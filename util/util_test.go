package util

import (
	v1 "configure/api/v1"
	"testing"
)

func TestProtoUnmarshal(t *testing.T) {

	pb := v1.AllEnvironmentReply{
		List: []*v1.AllEnvironmentReply_Environment{
			//{
			//	Keyword: "test1",
			//	Title:   "title",
			//},
		},
	}

	in := []struct {
		Keyword string
		Title   string
	}{
		{
			Keyword: "test1",
			Title:   "title",
		},
		{
			Keyword: "test2",
			Title:   "title2",
		},
	}

	if err := Transform(in, &pb.List); err != nil {
		t.Error(err.Error())
	}

	t.Log(in)
}
