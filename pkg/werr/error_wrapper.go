package werr

import "fmt"

func Wrap(context any, err any) error {
	if err == nil {
		if context == nil {
			return nil
		}
		return fmt.Errorf("%v", context)
	}

	switch context.(type) {
	case nil:
		return err.(error)
	case string:
		return fmt.Errorf("%s: %w", context, err.(error))
	case error:
		return fmt.Errorf("%w: %s", context.(error), err)
	default:
		return fmt.Errorf("%v: %w", context, err.(error))
	}
}
