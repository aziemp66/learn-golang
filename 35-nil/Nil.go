package main

func main() {

}

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	}
	return map[string]string{
		"name": name,
	}
}
