# ES基础查询参考

## 索引建立
```
POST nsf_data/_mapping
{
  "properties": {
    "award_effective_date": {
      "type ": "text",
      "fielddata": true
    }
  }
}

DELETE nsf_bucket/

PUT nsf_bucket/
{
  "settings": {
    "index": {
      "blocks.read_only_allow_delete":false
    }
  }, 
  "mappings" : {
      "properties" : {
        "@timestamp" : {
          "type" : "date"
        },
        "@version" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        },
        "XMLName" : {
          "properties" : {
            "Local" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            },
            "Space" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            }
          }
        },
        "award_amount" : {
          "type" : "long"
        },
        "award_effective_date" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        },
        "award_expiration_date" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        },
        "award_id" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        },
        "award_title" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        },
        "description" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        },
        "industries" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        },
        "institution" : {
          "properties" : {
            "cityname" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            },
            "country_name" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            },
            "name" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            },
            "phone_number" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            },
            "state_code" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            },
            "state_name" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            },
            "street_address" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            },
            "zipcode" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            }
          }
        },
        "investigator" : {
          "properties" : {
            "email_address" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            },
            "end_date" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            },
            "firstname" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            },
            "lastname" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            },
            "role_code" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            },
            "start_date" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            }
          }
        },
        "organization" : {
          "properties" : {
            "code" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            },
            "directorate" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            },
            "division" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            }
          }
        },
        "program_element" : {
          "properties" : {
            "code" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            },
            "text" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            }
          }
        },
        "program_reference" : {
          "properties" : {
            "code" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            },
            "text" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            }
          }
        },
        "technology" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        }
      }
    }
}

POST _reindex
{
  "source": {
    "index": "nsf_data"
  },
  "dest": {
    "index": "nsf_bucket"
  }
}

DELETE nsf_data

PUT nsf_data/
{
  "mappings" : {
      "properties" : {
        "award_amount" : {
          "type" : "long"
        },
        "award_effective_date" : {
          "type" : "date",
          "format": "MM/dd/yyyy"
        },
        "award_expiration_date" : {
          "type" : "date",
          "format": "MM/dd/yyyy"
        },
        "award_id" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        },
        "award_title" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        },
        "description" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        },
        "industries" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        },
        "technology" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        }
      }
    }
}

GET nsf_data/_mapping
```

GET nsf_data/_search
{
  "size": 20, 
  "query":{
    "match_all":{
    }
  }
}

## 主页

### 主页元素-排名-按金额排序
> GET /api/rank/search/amount

```
GET nsf_data/_search
{
  "size": 0,
  "aggs": {
    "amount_list": {
      "terms": {
        "field": "industries.keyword",
        "size": 2147483647
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
```

### 主页元素-排名-按增长率排序
> GET /api/rank/search/rate

```
GET nsf_data/_search
{
  "size": 0,
  "aggs": {
    "agg_year": {
      "date_histogram": {
        "field": "award_effective_date",
        "interval": "year",
        "format": "MM/dd/yyyy"
      },
      "aggs": {
        "rate_list": {
          "terms": {
            "field": "program_element.text.keyword",
            "size": 2147483647
          },
          "aggs": {
            "rate_sum": {
              "sum": {
                "field": "award_amount"
              }
            },
            "r_bucket_sort": {
              "bucket_sort": {
                "sort": {
                  "rate_sum": {
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
```

### 主页元素-解析-技术
> GET /api/tech/search/

```
GET nsf_data/_search
{
  "size": 0,
  "aggs": {
    "tech_group": {
      "terms": {
        "field": "program_element.text.keyword",
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
```

### 主页元素-解析-产业
>  GET /api/tech/search/

```
GET nsf_data/_search
{
  "size": 0,
  "aggs": {
    "indu_group": {
      "terms": {
        "field": "program_element.text.keyword",
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
```

### 主页元素-统计
> GET /api/basic/search/<br/>
#### fundingsAmount,基金总投资数

```
GET nsf_data/_search
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
```

#### fundingsInvested,基金平均投资数

```
GET nsf_data/_search
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
```

#### 热点产业

```
GET nsf_data/_search
{
  "size": 0,
  "aggs": {
    "hot_industry": {
      "terms": {
        "field": "industries.keyword",
        "size": 2147483647
      },
      "aggs": {
        "indu_sum": {
          "sum": {
            "field": "award_amount"
          }
        },
        "r_bucket_sort":{
          "bucket_sort": {
            "sort": {
              "indu_sum": {
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
```

#### 主分类资金分布饼图

```
GET nsf_data/_search
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
```

## 搜索页

### 搜索-简单搜索（没有使用）

```
GET nsf_data/_search
{
  "size": 20,
  "query": {
    "bool": {
      "should": [
        {
          "match": {
            "award_title":{
              "query": "Kinetic Theory",
              "operator": "and"
            }
          }
        },{
          "match": {
            "description": {
              "query": "significantly physical",
              "operator": "and"
            }
          }
        }
      ]
    }
  }
}
```

## 搜索-多项搜索
> POST /api/fund/search
```
GET nsf_data/_search
{
  "size": 20,
  "query": {
    "bool": {
      "should": [
        {
          "match": {
            "award_title":{
              "query": "Excellence Resource",
              "operator": "and"
            }
          }
        },{
          "match": {
            "description": {
              "query": "significantly physical",
              "operator": "and"
            }
          }
        },
        {
          "match": {
            "institution.name": {
              "query": "Ohio",
              "operator": "and"
            }
          }
        },
        {
          "match": {
            "organization.division": {
              "query": "Education",
              "operator": "and"
            }
          }
        },
        {
          "match": {
            "technology": {
              "query": "poly",
              "operator": "and"
            }
          }
        },
        {
          "match": {
            "industries": {
              "query": "design",
              "operator": "and"
            }
          }
        },
        {
          "range": {
            "award_effective_date":{
              "gte": "02/01/2019",
              "format": "MM/dd/yyyy||yyyy"
            }
          }
        },
        {
          "range": {
            "award_effective_date":{
              "lte": "09/30/2021",
              "format": "MM/dd/yyyy||yyyy"
            }
          }
        }
      ]
    }
  }
}
```

## 基金详情页

### 基金详情页-详情
> GET api/fund/detail/ + uuid

```
GET nsf_data/_search
{
  "size": 1,
  "query": {
    "term": {
      "award_id.keyword": {
        "value": "1900923"
      }
    }
  }
}
```

### 基金详情页-相关关联关系
> GET api/fund/cop<br>
> GET api/fund/word
#### 基金详情页-产业相关-(organization)机构（主分类）
```
GET nsf_data/_search
{
  "size": 0,
  "query": {
    "term": {
      "industries.keyword": {
        "value": "particle"
      }
    }
  },
  "aggs": {
    "related_division": {
      "terms": {
        "field": "organization.code.keyword",
        "size": 10
      },
      "aggs": {
        "division_name": {
          "terms": {
            "field": "organization.division.keyword",
            "size": 10
          },
          "aggs": {
            "proportion_fund": {
              "sum": {
                "field": "award_amount"
              }
            },
            "r_bucket_sort": {
              "bucket_sort": {
                "sort": {
                  "proportion_fund": {
                    "order": "desc"
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}
```

#### 基金详情页-产业相关-(institution)机构*
```
GET nsf_data/_search
{
  "size": 0,
  "query": {
    "term": {
      "industries.keyword": {
        "value": "particle"
      }
    }
  },
  "aggs": {
    "related_organizations": {
      "terms": {
        "field": "institution.name.keyword",
        "size": 10
      },
      "aggs": {
        "proportion_fund": {
          "sum": {
            "field": "award_amount"
          }
        },
        "r_bucket_sort": {
          "bucket_sort": {
            "sort": {
              "proportion_fund": {
                "order": "desc"
              }
            }
          }
        }
      }
    }
  }
}
```

#### 基金详情页-产业相关-技术
```
GET nsf_data/_search
{
  "size": 0,
  "query": {
    "term": {
      "industries.keyword": {
        "value": "particle"
      }
    }
  },
  "aggs": {
    "related_technologies": {
      "terms": {
        "field": "technology.keyword",
        "size": 10
      },
      "aggs": {
        "proportion_fund": {
          "sum": {
            "field": "award_amount"
          }
        },
        "r_bucket_sort": {
          "bucket_sort": {
            "sort": {
              "proportion_fund": {
                "order": "desc"
              }
            }
          }
        }
      }
    }
  }
}
```

#### 基金详情页-产业相关-主分类(废弃)
```
GET nsf_data/_search
{
  "size": 0,
  "query": {
    "term": {
      "industries.keyword": {
        "value": "particle"
      }
    }
  },
  "aggs": {
    "related_elements": {
      "terms": {
        "field": "program_element.code.keyword",
        "size": 10
      },
      "aggs": {
        "element_name": {
          "terms": {
            "field": "program_element.text.keyword",
            "size": 10
          },
          "aggs": {
            "proportion_fund": {
              "sum": {
                "field": "award_amount"
              }
            },
            "r_bucket_sort": {
              "bucket_sort": {
                "sort": {
                  "proportion_fund": {
                    "order": "desc"
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}
```

#### 基金详情页-技术相关-主分类
```
GET nsf_data/_search
{
  "size": 0,
  "query": {
    "term": {
      "technology.keyword": {
        "value": "magnetic field"
      }
    }
  },
  "aggs": {
    "related_division": {
      "terms": {
        "field": "organization.code.keyword",
        "size": 10
      },
      "aggs": {
        "division_name": {
          "terms": {
            "field": "organization.division.keyword",
            "size": 10
          },
          "aggs": {
            "proportion_fund": {
              "sum": {
                "field": "award_amount"
              }
            },
            "r_bucket_sort": {
              "bucket_sort": {
                "sort": {
                  "proportion_fund": {
                    "order": "desc"
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}
```

#### 基金详情页-技术相关-产业
```
GET nsf_data/_search
{
  "size": 0,
  "query": {
    "term": {
      "technology.keyword": {
        "value": "magnetic field"
      }
    }
  },
  "aggs": {
    "related_industries": {
      "terms": {
        "field": "industries.keyword",
        "size": 10
      },
      "aggs": {
        "proportion_fund": {
          "sum": {
            "field": "award_amount"
          }
        },
        "r_bucket_sort": {
          "bucket_sort": {
            "sort": {
              "proportion_fund": {
                "order": "desc"
              }
            }
          }
        }
      }
    }
  }
}
```

#### 基金详情页-技术相关-机构
```
GET nsf_data/_search
{
  "size": 0,
  "query": {
    "term": {
      "technology.keyword": {
        "value": "magnetic field"
      }
    }
  },
  "aggs": {
    "related_organizations": {
      "terms": {
        "field": "institution.name.keyword",
        "size": 10
      },
      "aggs": {
        "proportion_fund": {
          "sum": {
            "field": "award_amount"
          }
        },
        "r_bucket_sort": {
          "bucket_sort": {
            "sort": {
              "proportion_fund": {
                "order": "desc"
              }
            }
          }
        }
      }
    }
  }
}
```

#### 基金详情页-主分类相关-产业
```
GET nsf_data/_search
{
  "size": 0,
  "query": {
    "term": {
      "organization.code.keyword": {
        "value": "03010000"
      }
    }
  },
  "aggs": {
    "related_industries": {
      "terms": {
        "field": "industries.keyword",
        "size": 10
      },
      "aggs": {
        "proportion_fund": {
          "sum": {
            "field": "award_amount"
          }
        },
        "r_bucket_sort": {
          "bucket_sort": {
            "sort": {
              "proportion_fund": {
                "order": "desc"
              }
            }
          }
        }
      }
    }
  }
}
```

#### 基金详情页-主分类相关-技术
```
GET nsf_data/_search
{
  "size": 0,
  "query": {
    "term": {
      "organization.code.keyword": {
        "value": "03010000"
      }
    }
  },
  "aggs": {
    "related_technologies": {
      "terms": {
        "field": "technology.keyword",
        "size": 10
      },
      "aggs": {
        "proportion_fund": {
          "sum": {
            "field": "award_amount"
          }
        },
        "r_bucket_sort": {
          "bucket_sort": {
            "sort": {
              "proportion_fund": {
                "order": "desc"
              }
            }
          }
        }
      }
    }
  }
}
```

#### 基金详情页-主分类相关-机构
```
GET nsf_data/_search
{
  "size": 0,
  "query": {
    "term": {
      "organization.code.keyword": {
        "value": "03010000"
      }
    }
  },
  "aggs": {
    "related_organizations": {
      "terms": {
        "field": "institution.name.keyword",
        "size": 10
      },
      "aggs": {
        "proportion_fund": {
          "sum": {
            "field": "award_amount"
          }
        },
        "r_bucket_sort": {
          "bucket_sort": {
            "sort": {
              "proportion_fund": {
                "order": "desc"
              }
            }
          }
        }
      }
    }
  }
}
```

#### 基金详情页-机构相关-产业
```
GET nsf_data/_search
{
  "size": 0,
  "query": {
    "term": {
      "institution.name.keyword": {
        "value": "Northwestern University"
      }
    }
  },
  "aggs": {
    "related_industries": {
      "terms": {
        "field": "industries.keyword",
        "size": 10
      },
      "aggs": {
        "proportion_fund": {
          "sum": {
            "field": "award_amount"
          }
        },
        "r_bucket_sort": {
          "bucket_sort": {
            "sort": {
              "proportion_fund": {
                "order": "desc"
              }
            }
          }
        }
      }
    }
  }
}
```

#### 基金详情页-机构相关-技术
```
GET nsf_data/_search
{
  "size": 0,
  "query": {
    "term": {
      "institution.name.keyword": {
        "value": "Northwestern University"
      }
    }
  },
  "aggs": {
    "related_technologies": {
      "terms": {
        "field": "technology.keyword",
        "size": 10
      },
      "aggs": {
        "proportion_fund": {
          "sum": {
            "field": "award_amount"
          }
        },
        "r_bucket_sort": {
          "bucket_sort": {
            "sort": {
              "proportion_fund": {
                "order": "desc"
              }
            }
          }
        }
      }
    }
  }
}
```

#### 基金详情页-机构相关-主分类
```
GET nsf_data/_search
{
  "size": 0,
  "query": {
    "term": {
      "institution.name.keyword": {
        "value": "Northwestern University"
      }
    }
  },
  "aggs": {
    "related_division": {
      "terms": {
        "field": "organization.code.keyword",
        "size": 10
      },
      "aggs": {
        "division_name": {
          "terms": {
            "field": "organization.division.keyword",
            "size": 10
          },
          "aggs": {
            "proportion_fund": {
              "sum": {
                "field": "award_amount"
              }
            },
            "r_bucket_sort": {
              "bucket_sort": {
                "sort": {
                  "proportion_fund": {
                    "order": "desc"
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}
```

### 基金详情页-关联基金
> 根据相关分析的结果，取基金相关产业、技术、主分类、机构（top 5）的基金号的并集，为关联基金

> GET /api/fund/similar
```
GET nsf_data/_search
{
  "size": 10,
  "query": {
    "term": {
      "organization.code.keyword": {
        "value": "03010000"
      }
    }
  }
}
```
## 主分类详情页

### 主分类详情页-基金投资金额排名
```
GET nsf_data/_search
{
  "size": 20,
  "query": {
    "term": {
      "organization.code.keyword": {
        "value": "03010000"
      }
    }
  },
  "sort": [
    {
      "award_amount": {
        "order": "desc"
      }
    }
  ]
}
```

### 主分类详情页-技术投资排名
```
GET nsf_data/_search
{
  "size": 0,
  "query": {
    "term": {
      "organization.code.keyword": {
        "value": "03010000"
      }
    }
  },
  "aggs": {
    "tech_list": {
      "terms": {
        "field": "technology.keyword",
        "size": 10
      },
      "aggs": {
        "year_bucket": {
          "date_histogram": {
            "field": "award_effective_date",
            "interval": "year"
          },
          "aggs": {
            "proportion_fund": {
              "sum": {
                "field": "award_amount"
              }
            },
            "r_bucket_sort": {
              "bucket_sort": {
                "sort": {
                  "proportion_fund": {
                    "order": "desc"
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}
```

### 主分类详情页-产业投资排名
```
GET nsf_data/_search
{
  "size": 0,
  "query": {
    "term": {
      "organization.code.keyword": {
        "value": "03010000"
      }
    }
  },
  "aggs": {
    "tech_list": {
      "terms": {
        "field": "industries.keyword",
        "size": 10
      },
      "aggs": {
        "year_bucket": {
          "date_histogram": {
            "field": "award_effective_date",
            "interval": "year"
          },
          "aggs": {
            "proportion_fund": {
              "sum": {
                "field": "award_amount"
              }
            },
            "r_bucket_sort": {
              "bucket_sort": {
                "sort": {
                  "proportion_fund": {
                    "order": "desc"
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}
```

### 主分类详情页-机构投资排名
```
GET nsf_data/_search
{
  "size": 0,
  "query": {
    "term": {
      "organization.code.keyword": {
        "value": "03010000"
      }
    }
  },
  "aggs": {
    "tech_list": {
      "terms": {
        "field": "institution.name.keyword",
        "size": 10
      },
      "aggs": {
        "year_bucket": {
          "date_histogram": {
            "field": "award_effective_date",
            "interval": "year"
          },
          "aggs": {
            "proportion_fund": {
              "sum": {
                "field": "award_amount"
              }
            },
            "r_bucket_sort": {
              "bucket_sort": {
                "sort": {
                  "proportion_fund": {
                    "order": "desc"
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}
```

## 机构详情页

### 机构详情页-基金投资金额排名
```
GET nsf_data/_search
{
  "size": 20,
  "query": {
    "term": {
      "institution.name.keyword": {
        "value": "Princeton University"
      }
    }
  },
  "sort": [
    {
      "award_amount": {
        "order": "desc"
      }
    }
  ]
}
```

### 机构详情页-技术投资排名
```
GET nsf_data/_search
{
  "size": 0,
  "query": {
    "term": {
      "institution.name.keyword": {
        "value": "Princeton University"
      }
    }
  },
  "aggs": {
    "tech_list": {
      "terms": {
        "field": "technology.keyword",
        "size": 10
      },
      "aggs": {
        "year_bucket": {
          "date_histogram": {
            "field": "award_effective_date",
            "interval": "year"
          },
          "aggs": {
            "proportion_fund": {
              "sum": {
                "field": "award_amount"
              }
            },
            "r_bucket_sort": {
              "bucket_sort": {
                "sort": {
                  "proportion_fund": {
                    "order": "desc"
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}
```

### 机构详情页-产业投资排名
```
GET nsf_data/_search
{
  "size": 0,
  "query": {
    "term": {
      "institution.name.keyword": {
        "value": "Princeton University"
      }
    }
  },
  "aggs": {
    "tech_list": {
      "terms": {
        "field": "industries.keyword",
        "size": 10
      },
      "aggs": {
        "year_bucket": {
          "date_histogram": {
            "field": "award_effective_date",
            "interval": "year"
          },
          "aggs": {
            "proportion_fund": {
              "sum": {
                "field": "award_amount"
              }
            },
            "r_bucket_sort": {
              "bucket_sort": {
                "sort": {
                  "proportion_fund": {
                    "order": "desc"
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}
```

### 机构详情页-主分类投资排名
```
GET nsf_data/_search
{
  "size": 0,
  "query": {
    "term": {
      "institution.name.keyword": {
        "value": "Princeton University"
      }
    }
  },
  "aggs": {
    "division_list": {
      "terms": {
        "field": "organization.code.keyword",
        "size": 10
      },
      "aggs": {
        "division_name": {
          "terms": {
            "field": "organization.division.keyword",
            "size": 10
          },
          "aggs": {
            "year_bucket": {
              "date_histogram": {
                "field": "award_effective_date",
                "interval": "year"
              },
              "aggs": {
                "proportion_fund": {
                  "sum": {
                    "field": "award_amount"
                  }
                },
                "r_bucket_sort": {
                  "bucket_sort": {
                    "sort": {
                      "proportion_fund": {
                        "order": "desc"
                      }
                    }
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}
```

## 技术详情页

### 技术详情页-基金投资金额排名
```
GET nsf_data/_search
{
  "size": 20,
  "query": {
    "term": {
      "technology.keyword": {
        "value": "automorphic form"
      }
    }
  },
  "sort": [
    {
      "award_amount": {
        "order": "desc"
      }
    }
  ]
}
```

### 技术详情页-机构投资排名
```
GET nsf_data/_search
{
  "size": 0,
  "query": {
    "term": {
      "technology.keyword": {
        "value": "automorphic form"
      }
    }
  },
  "aggs": {
    "tech_list": {
      "terms": {
        "field": "institution.name.keyword",
        "size": 10
      },
      "aggs": {
        "year_bucket": {
          "date_histogram": {
            "field": "award_effective_date",
            "interval": "year"
          },
          "aggs": {
            "proportion_fund": {
              "sum": {
                "field": "award_amount"
              }
            },
            "r_bucket_sort": {
              "bucket_sort": {
                "sort": {
                  "proportion_fund": {
                    "order": "desc"
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}
```

### 技术详情页-产业投资排名
```
GET nsf_data/_search
{
  "size": 0,
  "query": {
    "term": {
      "technology.keyword": {
        "value": "automorphic form"
      }
    }
  },
  "aggs": {
    "tech_list": {
      "terms": {
        "field": "industries.keyword",
        "size": 10
      },
      "aggs": {
        "year_bucket": {
          "date_histogram": {
            "field": "award_effective_date",
            "interval": "year"
          },
          "aggs": {
            "proportion_fund": {
              "sum": {
                "field": "award_amount"
              }
            },
            "r_bucket_sort": {
              "bucket_sort": {
                "sort": {
                  "proportion_fund": {
                    "order": "desc"
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}
```

### 技术详情页-主分类投资排名
```
GET nsf_data/_search
{
  "size": 0,
  "query": {
    "term": {
      "technology.keyword": {
        "value": "automorphic form"
      }
    }
  },
  "aggs": {
    "division_list": {
      "terms": {
        "field": "organization.code.keyword",
        "size": 10
      },
      "aggs": {
        "division_name": {
          "terms": {
            "field": "organization.division.keyword",
            "size": 10
          },
          "aggs": {
            "year_bucket": {
              "date_histogram": {
                "field": "award_effective_date",
                "interval": "year"
              },
              "aggs": {
                "proportion_fund": {
                  "sum": {
                    "field": "award_amount"
                  }
                },
                "r_bucket_sort": {
                  "bucket_sort": {
                    "sort": {
                      "proportion_fund": {
                        "order": "desc"
                      }
                    }
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}
```

## 产业详情页
### 产业详情页-基金投资金额排名
```
GET nsf_data/_search
{
  "size": 20,
  "query": {
    "term": {
      "industries.keyword": {
        "value": "number theory"
      }
    }
  },
  "sort": [
    {
      "award_amount": {
        "order": "desc"
      }
    }
  ]
}
```

### 产业详情页-机构投资排名
```
GET nsf_data/_search
{
  "size": 0,
  "query": {
    "term": {
      "industries.keyword": {
        "value": "number theory"
      }
    }
  },
  "aggs": {
    "tech_list": {
      "terms": {
        "field": "institution.name.keyword",
        "size": 10
      },
      "aggs": {
        "year_bucket": {
          "date_histogram": {
            "field": "award_effective_date",
            "interval": "year"
          },
          "aggs": {
            "proportion_fund": {
              "sum": {
                "field": "award_amount"
              }
            },
            "r_bucket_sort": {
              "bucket_sort": {
                "sort": {
                  "proportion_fund": {
                    "order": "desc"
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}
```

### 产业详情页-技术投资排名
```
GET nsf_data/_search
{
  "size": 0,
  "query": {
    "term": {
      "industries.keyword": {
        "value": "number theory"
      }
    }
  },
  "aggs": {
    "tech_list": {
      "terms": {
        "field": "technology.keyword",
        "size": 10
      },
      "aggs": {
        "year_bucket": {
          "date_histogram": {
            "field": "award_effective_date",
            "interval": "year"
          },
          "aggs": {
            "proportion_fund": {
              "sum": {
                "field": "award_amount"
              }
            },
            "r_bucket_sort": {
              "bucket_sort": {
                "sort": {
                  "proportion_fund": {
                    "order": "desc"
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}
```

### 产业详情页-主分类投资排名
```
GET nsf_data/_search
{
  "size": 0,
  "query": {
    "term": {
      "industries.keyword": {
        "value": "number theory"
      }
    }
  },
  "aggs": {
    "division_list": {
      "terms": {
        "field": "organization.code.keyword",
        "size": 10
      },
      "aggs": {
        "division_name": {
          "terms": {
            "field": "organization.division.keyword",
            "size": 10
          },
          "aggs": {
            "year_bucket": {
              "date_histogram": {
                "field": "award_effective_date",
                "interval": "year"
              },
              "aggs": {
                "proportion_fund": {
                  "sum": {
                    "field": "award_amount"
                  }
                },
                "r_bucket_sort": {
                  "bucket_sort": {
                    "sort": {
                      "proportion_fund": {
                        "order": "desc"
                      }
                    }
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}
```


## Additions

### fundRank投资排名
```
GET nsf_data/_search
{
  "size": 20,
  "query": {
    "match_all": {}
  },
  "sort": [
    {
      "award_amount": {
        "order": "desc"
      }
    }
  ]
}
```

### 投资排名翻页
```
GET nsf_data/_search
{
  "size": 20,
  "query": {
    "match_all": {}
  },
  "search_after":[449999],
  "sort": [
    {
      "award_amount": {
        "order": "desc"
      }
    }
  ]
}
```

```
GET nsf_test/_search
{
  "query": {
    "term": {
      "doc.Award.AwardID": {
        "value": "1903568"
      }
    }
  }
}
```
