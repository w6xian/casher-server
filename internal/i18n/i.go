package i18n

type IParse interface {
	// String 获取翻译后的字符串
	L(key, def string, args ...Field) string
}
