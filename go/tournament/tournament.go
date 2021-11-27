package tournament

import "io"
import "fmt"
import "errors"
import "strings"
import "bytes"
import "sort"
import "text/tabwriter"

const bToRead = 100

type ScoreRecord struct {
	Team string
	MP   int
	W    int
	D    int
	L    int
	P    int
}

// Scoreboard type
type ScoreBType map[string]*ScoreRecord

// sort by Team Name (Descending)
type byTeamScoreB []*ScoreRecord

func (s byTeamScoreB) Len() int           { return len(s) }
func (s byTeamScoreB) Less(i, j int) bool { return s[i].Team < s[j].Team }
func (s byTeamScoreB) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// sort by Point (Descending)
type byPointScoreB []*ScoreRecord

func (s byPointScoreB) Len() int           { return len(s) }
func (s byPointScoreB) Less(i, j int) bool { return s[i].P > s[j].P }
func (s byPointScoreB) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// Score Board
var ScoreB ScoreBType
var sb []*ScoreRecord

// Tally parse text and outpup Scoreboard
func Tally(reader io.Reader, writer io.Writer) error {
	ScoreB = ScoreBType{}
	sb = make([]*ScoreRecord, 0, 10)
	bchunk := make([]byte, bToRead)
	var btext bytes.Buffer

	// get text
	for {
		nbytes, err := reader.Read(bchunk)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		btext.Write(bchunk[:nbytes])
	}

	lineSlice := strings.Split(strings.Trim(btext.String(), " "), "\n")

	for _, l := range lineSlice {

		splited := strings.Split(l, ";")

		if len(splited) < 2 {
			continue
		} else if len(splited) < 3 {
			return errors.New("incorrect input.")
		}

		err := fillScoreB(splited[0], splited[1], splited[2])
		if err != nil {
			return err
		}
	}
	// sort score slice
	sort.Sort(byTeamScoreB(sb))
	sort.Sort(byPointScoreB(sb))

	_, err := writer.Write(outputScoreB())
	if err != nil {
		return err
	}

	return nil
}

// outputScoreB formats byte slice output
func outputScoreB() []byte {
	var r bytes.Buffer
	format := "%v\t\t\t\t| %v |\t%v |\t%v |\t%v |\t%v\n"
	tw := tabwriter.NewWriter(&r, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Team", "MP", "W", "D", "L", "P")
	format = "%v\t\t\t\t|  %v |\t%v |\t%v |\t%v |\t%v\n"
	for _, v := range sb {
		fmt.Fprintf(tw, format, v.Team, v.MP, v.W, v.D, v.L, v.P)
	}
	tw.Flush()
	return r.Bytes()
}

// fillScoreB parses input text and fills scoreboard
func fillScoreB(team1, team2, res string) error {
	team1 = strings.Trim(team1, " ")
	team2 = strings.Trim(team2, " ")
	res = strings.Trim(res, " ")

	t1, ok := ScoreB[team1]
	if !ok {
		t1 = &ScoreRecord{Team: team1}
		ScoreB[team1] = t1
		sb = append(sb, t1)
	}
	t1.MP += 1

	t2, ok := ScoreB[team2]
	if !ok {
		t2 = &ScoreRecord{Team: team2}
		ScoreB[team2] = t2
		sb = append(sb, t2)
	}
	t2.MP += 1

	switch res {
	case "win":
		t1.W += 1
		t1.P += 3
		t2.L += 1
	case "loss":
		t1.L += 1
		t2.W += 1
		t2.P += 3
	case "draw":
		t1.D += 1
		t1.P += 1
		t2.D += 1
		t2.P += 1
	default:
		return errors.New("incorrrect input.")
	}
	return nil
}
