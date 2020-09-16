package dbops

import "fmt"

const (
	organization = "institution.name.keyword"
	technology   = "technology.keyword"
	division     = "organization.code.keyword"
	industry     = "industries.keyword"
)

var basic_analysis_tech = `
{
  "size": 0,
  "aggs": {
    "tech_group": {
      "terms": {
        "field": "organization.division.keyword",
        "size": 100
      },
      "aggs": {
        "techs": {
          "terms": {
            "field": "technology.keyword",
            "size": 100
          },
          "aggs": {
            "tech_group_sum": {
              "sum": {
                "field": "award_amount"
              }
            },
            "r_bucket_sort":{
              "bucket_sort": {
                "sort": {
                  "tech_group_sum": {
                    "order": "desc"
                  }
                },
                "size": 10
              }
            }
          }
        }
      }
    }
  }
}
`

var basic_analysis_indu = `
{
	"size": 0,
	"aggs": {
	  "indu_group": {
		"terms": {
		  "field": "organization.division.keyword",
		  "size": 50
		},
		"aggs": {
		  "indus": {
			"terms": {
			  "field": "industries.keyword",
			  "size": 50
			},
			"aggs": {
			  "indu_group_sum": {
				"sum": {
				  "field": "award_amount"
				}
			  },
			  "r_bucket_sort":{
				"bucket_sort": {
				  "sort": {
					"indu_group_sum": {
					  "order": "desc"
					}
				  },
				  "size": 10
				}
			  }
			}
		  }
		}
	  }
	}
}
`

var basic_statistics_amount = `
{
    "size":0,
    "aggs":{
      "fundings_amount": {
        "sum": {
          "field":"award_amount"
        }
      }
    }
  }
`

var basic_statistics_avg = `
{
    "size": 0,
    "aggs": {
      "avg_invested": {
        "avg": {
          "field": "award_amount"
        }
      }
    }
  }
`

var basic_pie = `
{
    "size": 0,
    "aggs": {
      "categorys": {
        "terms": {
          "field": "organization.code.keyword",
          "size": 2147483647
        },
        "aggs": {
          "category_name": {
            "terms": {
              "field": "organization.division.keyword",
              "size": 10
            },
            "aggs": {
              "category_proportion": {
                "sum": {
                  "field": "award_amount"
                }
              },
              "r_bucket_sort":{
                "bucket_sort": {
                  "sort": {
                    "category_proportion": {
                      "order": "desc"
                    }
                  },
                  "size": 10
                }
              }
            }
          }
        }
      }
    }
  }
`

func GetRelatedTemplate(key, value string) string {
	var basic_related = `
    {
        "size": 10,
        "query": {
            "term": {
                "%s": {
                    "value": %s
                }
            }
        }
    }
    `
	return fmt.Sprintf(basic_related, key, value)
}
