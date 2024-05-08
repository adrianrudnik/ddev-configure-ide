package xmlknife

import "github.com/beevik/etree"

func OverwriteXmlElement(e *etree.Element, tag string, value string) {
	if fe := e.FindElement(tag); fe != nil {
		fe.SetText(value)
	} else {
		e.CreateElement(tag).SetText(value)
	}
}
