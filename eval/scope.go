package eval

const (
	_ = iota
	GlobalScope
	FuncScope
	IFScope
)

type ScopeType int

// 变量作用域管理
type ScopeManager struct {
	curScope    *Scope
	globalScope *Scope
}

func NewScopeManager() *ScopeManager {
	sm := &ScopeManager{}
	globalScope := &Scope{
		Container: make(map[string]interface{}),
		Parent:    nil,
		Type:      GlobalScope,
	}
	sm.curScope = globalScope
	sm.globalScope = globalScope
	return sm
}

func (sm *ScopeManager) Push(scopeType ScopeType) {
	newScope := &Scope{
		Container: make(map[string]interface{}),
		Parent:    sm.curScope,
		Type:      scopeType,
	}
	sm.curScope = newScope
}

func (sm *ScopeManager) Pop() {
	parent := sm.curScope.Parent
	if parent.Type == GlobalScope {
		return
	}
	sm.curScope = parent
}

func (sm *ScopeManager) GetValue(name string) (interface{}, *Scope) {
	curScopeType := sm.curScope.Type
	for scope := sm.curScope; scope != nil; {
		v, ok := scope.Container[name]
		if ok {
			return v, scope
		}
		switch curScopeType {
		case GlobalScope:
			return nil, nil

		case FuncScope:
			scope = sm.globalScope

		case IFScope:
			scope = scope.Parent

		default:
			panic("未知的作用域")
		}

	}
	return nil, nil
}

func (sm *ScopeManager) SetValue(name string, value interface{}, newVar bool) bool {
	if newVar {
		sm.curScope.Container[name] = value
		return true
	}

	if !newVar {
		_, scope := sm.GetValue(name)
		if scope == nil {
			return false
		}
		scope.Container[name] = value
		return true
	}
	return true
}

func (sm *ScopeManager) VarExists(name string) bool {
	_, scope := sm.GetValue(name)
	return scope != nil
}

type Scope struct {
	Container map[string]interface{}
	Parent    *Scope
	Type      ScopeType // 当前作用域类型  global function control
}
