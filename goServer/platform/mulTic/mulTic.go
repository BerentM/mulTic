package mulTic

type Status struct {
	NextPlayer string `json:"nextPlayer"`
	Board      string `json:"board"`
	Winner     string `json:"winner"`
}

func New() *Status {
	return &Status{}
}

func (r *Status) Update(b string) {
	st := r.Show()
	if st.NextPlayer == "x" {
		r.NextPlayer = "o"
	} else {
		r.NextPlayer = "x"
	}
	r.Board = b
	r.Winner = ""
}

func (r *Status) Restart(n string) {
	r.NextPlayer = n
	r.Board = `{"0": " ", "1": " ", "2": " ", "3": " ", "4": " ", "5": " ", "6": " ", "7": " ", "8": " "}`
	r.Winner = ""
}

func (r *Status) Show() Status {
	return *r
}
