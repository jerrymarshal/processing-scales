func parseGadgets(path string, s string) ([]Gadget, error) {
	gadgets := []Gadget{}
	const winNewline = "\r\n"
	const normalNewline = "\n"
	nl := normalNewline
	if strings.Contains(s, winNewline) {
		nl = winNewline
	}


  func ExpandGadget(gadget *Gadget, params []string) string {
	str := gadget.json
	for i := 0; i < gadget.parametersCount; i++ {
		str = strings.ReplaceAll(str, "%"+strconv.Itoa(i), strings.TrimSpace(params[i]))
	}
	return str

	  var dir string
	// we will continue searching thru the tree until we return
	for {
		dir, _, gadgets, err = loadGadgetsForPath(searchDir, gadgets)
		if err != nil || dir == searchDir {
			return gadgets, err
		}
}

  //taking a sample
