package mulTic

type Status struct {
	NextPlayer string `json:"nextPlayer"`
	Board      string `json:"board"`
	Winner     string `json:"winner"`
}

func New() *Status {
	return &Status{}
}

func (r *Status) Update(s string) {
	st := r.Show()
	if st.NextPlayer == "x" {
		r.NextPlayer = "o"
	} else {
		r.NextPlayer = "x"
	}
	r.Board = s
	r.Winner = ""
}

func (r *Status) Show() Status {
	return *r
}
