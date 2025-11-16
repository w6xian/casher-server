package checker

import "fmt"

type CheckType int

const (
	CHECKER_TYPE_INT CheckType = iota
	// int32
	CHECKER_TYPE_INT32
	// int64
	CHECKER_TYPE_INT64
	// uint
	CHECKER_TYPE_UINT
	// uint32
	CHECKER_TYPE_UINT32
	// uint64
	CHECKER_TYPE_UINT64
	// float32
	CHECKER_TYPE_FLOAT32
	// float64
	CHECKER_TYPE_FLOAT64
	CHECKER_TYPE_STRING
	// bool
	CHECKER_TYPE_BOOL
	CHECKER_TYPE_OTHER
)

type Checks struct {
	Values map[string]*Checker
}

func New(checkers ...*Checker) *Checks {
	c := &Checks{
		Values: make(map[string]*Checker),
	}
	c.Add(checkers...)
	return c
}
func (c *Checks) Add(checkers ...*Checker) {
	// for checkers
	for _, checker := range checkers {
		c.Values[checker.Name] = checker
	}
}
func (c *Checks) Get(name string) *Checker {
	if checker, ok := c.Values[name]; ok {
		return checker
	}
	return nil
}

func (c *Checks) CheckMap(values map[string]any) (map[string]any, error) {
	kv := make(map[string]any)
	// for values check
	for k, v := range values {
		checker := c.Get(k)
		if checker == nil {
			continue
		}
		rv, err := checker.Check(v)
		if err != nil {
			return nil, err
		}
		kv[k] = rv
	}
	return kv, nil
}
func (c *Checks) Check(values map[string]any) error {
	// for values check
	for k, v := range values {
		checker := c.Get(k)
		if checker == nil {
			continue
		}
		_, err := checker.Check(v)
		if err != nil {
			return err
		}
	}
	return nil
}

type Checker struct {
	Name string
	Type CheckType
	opts Options
}

// check value
func (c *Checker) Check(value any) (any, error) {
	switch c.Type {
	case CHECKER_TYPE_INT:
		v, ok := value.(int)
		if !ok {
			return nil, fmt.Errorf("checker %s type error, expect int, got %T", c.Name, value)
		}
		return v, nil
	case CHECKER_TYPE_INT32:
		v, ok := value.(int32)
		if !ok {
			return nil, fmt.Errorf("checker %s type error, expect int32, got %T", c.Name, value)
		}
		return v, nil
	case CHECKER_TYPE_INT64:
		v, ok := value.(int64)
		if !ok {
			return nil, fmt.Errorf("checker %s type error, expect int64, got %T", c.Name, value)
		}
		return v, nil
	case CHECKER_TYPE_UINT:
		v, ok := value.(uint)
		if !ok {
			return nil, fmt.Errorf("checker %s type error, expect uint, got %T", c.Name, value)
		}
		return v, nil
	case CHECKER_TYPE_UINT32:
		v, ok := value.(uint32)
		if !ok {
			return nil, fmt.Errorf("checker %s type error, expect uint32, got %T", c.Name, value)
		}
		return v, nil
	case CHECKER_TYPE_UINT64:
		v, ok := value.(uint64)
		if !ok {
			return nil, fmt.Errorf("checker %s type error, expect uint64, got %T", c.Name, value)
		}
		return v, nil
	case CHECKER_TYPE_FLOAT32:
		v, ok := value.(float32)
		if !ok {
			return nil, fmt.Errorf("checker %s type error, expect float32, got %T", c.Name, value)
		}
		return v, nil
	case CHECKER_TYPE_FLOAT64:
		v, ok := value.(float64)
		if !ok {
			return nil, fmt.Errorf("checker %s type error, expect float64, got %T", c.Name, value)
		}
		return v, nil
	case CHECKER_TYPE_STRING:
		v, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("checker %s type error, expect string, got %T", c.Name, value)
		}
		return v, nil
	case CHECKER_TYPE_BOOL:
		v, ok := value.(bool)
		if !ok {
			return nil, fmt.Errorf("checker %s type error, expect bool, got %T", c.Name, value)
		}
		return v, nil
	case CHECKER_TYPE_OTHER:
		if c.opts.checker == nil {
			return nil, fmt.Errorf("checker %s type error, expect func, got nil", c.Name)
		}
		return c.opts.checker(value)
	default:
		return nil, fmt.Errorf("checker %s type error, unknown type %d", c.Name, c.Type)
	}
}

// check
