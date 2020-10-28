package mulTic

import (
	"encoding/json"
)

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
	var m map[string]interface{}
	json.Unmarshal([]byte(b), &m)

	win := ((m["0"] == st.NextPlayer && m["1"] == st.NextPlayer && m["2"] == st.NextPlayer) ||
		(m["3"] == st.NextPlayer && m["4"] == st.NextPlayer && m["5"] == st.NextPlayer) ||
		(m["6"] == st.NextPlayer && m["7"] == st.NextPlayer && m["8"] == st.NextPlayer) ||
		(m["0"] == st.NextPlayer && m["4"] == st.NextPlayer && m["8"] == st.NextPlayer) ||
		(m["2"] == st.NextPlayer && m["4"] == st.NextPlayer && m["6"] == st.NextPlayer) ||
		(m["0"] == st.NextPlayer && m["3"] == st.NextPlayer && m["6"] == st.NextPlayer) ||
		(m["1"] == st.NextPlayer && m["4"] == st.NextPlayer && m["7"] == st.NextPlayer) ||
		(m["2"] == st.NextPlayer && m["5"] == st.NextPlayer && m["8"] == st.NextPlayer))

	if win {
		r.Winner = st.NextPlayer
	}

	if st.NextPlayer == "x" {
		r.NextPlayer = "o"
	} else {
		r.NextPlayer = "x"
	}
	r.Board = b
}

func (r *Status) Restart(n string) {
	r.NextPlayer = n
	r.Board = `{"0": " ", "1": " ", "2": " ", "3": " ", "4": " ", "5": " ", "6": " ", "7": " ", "8": " "}`
	r.Winner = ""
}

func (r *Status) Show() Status {
	return *r
}
