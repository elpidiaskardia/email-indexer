package models

type ZincsearchResult  struct {
    Hits struct {
        Total struct {
            Value int `json:"value"`
        } `json:"total"`
        Hits []struct {
            Source EmailData `json:"_source"`
        } `json:"hits"`
    } `json:"hits"`
}