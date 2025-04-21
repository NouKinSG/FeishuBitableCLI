package tui

// displayOr 返回 val（非空）或 def
func displayOr(val, def string) string {
	if val == "" {
		return def
	}
	return val
}
