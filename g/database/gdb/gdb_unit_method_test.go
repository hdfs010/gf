// 方法操作

package gdb_test

import (
    "github.com/gogf/gf/g"
    "github.com/gogf/gf/g/os/gtime"
    "github.com/gogf/gf/g/test/gtest"
    "testing"
)

func TestDbBase_Query(t *testing.T) {
    if _, err := db.Query("SELECT ?", 1); err != nil {
        gtest.Fatal(err)
    }
    if _, err := db.Query("ERROR"); err == nil {
        gtest.Fatal("FAIL")
    }
}

func TestDbBase_Exec(t *testing.T) {
	if _, err := db.Exec("SELECT ?", 1); err != nil {
		gtest.Fatal(err)
	}
	if _, err := db.Exec("ERROR"); err == nil {
		gtest.Fatal("FAIL")
	}
}

func TestDbBase_Prepare(t *testing.T) {
    st, err := db.Prepare("SELECT 100")
    if err != nil {
        gtest.Fatal(err)
    }
    rows, err := st.Query()
    if err != nil {
        gtest.Fatal(err)
    }
    array, err := rows.Columns()
    if err != nil {
        gtest.Fatal(err)
    }
    gtest.Assert(array[0], "100")
    if err := rows.Close(); err != nil {
        gtest.Fatal(err)
    }
}

func TestDbBase_Insert(t *testing.T) {
    if _, err := db.Insert("user", g.Map{
        "id"          : 1,
        "passport"    : "t1",
        "password"    : "25d55ad283aa400af464c76d713c07ad",
        "nickname"    : "T1",
        "create_time" : gtime.Now().String(),
    }); err != nil {
        gtest.Fatal(err)
    }
    // normal map
    result, err := db.Insert("user", map[interface{}]interface{} {
        "id"          : "2",
        "passport"    : "t2",
        "password"    : "25d55ad283aa400af464c76d713c07ad",
        "nickname"    : "T2",
        "create_time" : gtime.Now().String(),
    })
    if err != nil {
        gtest.Fatal(err)
    }
    n, _ := result.RowsAffected()
    gtest.Assert(n, 1)

    // struct
    type User struct {
        Id         int    `gconv:"id"`
        Passport   string `json:"passport"`
        Password   string `gconv:"password"`
        Nickname   string `gconv:"nickname"`
        CreateTime string `json:"create_time"`
    }
    result, err = db.Insert("user", User{
        Id          : 3,
        Passport    : "t3",
        Password    : "25d55ad283aa400af464c76d713c07ad",
        Nickname    : "T3",
        CreateTime  : gtime.Now().String(),
    })
    if err != nil {
        gtest.Fatal(err)
    }
    n, _ = result.RowsAffected()
    gtest.Assert(n, 1)
    value, err := db.GetValue("select `passport` from `user` where id=?", 3)
    gtest.Assert(err, nil)
    gtest.Assert(value.String(), "t3")

    // *struct
    result, err = db.Insert("user", &User{
        Id          : 4,
        Passport    : "t4",
        Password    : "25d55ad283aa400af464c76d713c07ad",
        Nickname    : "T4",
        CreateTime  : gtime.Now().String(),
    })
    if err != nil {
        gtest.Fatal(err)
    }
    n, _ = result.RowsAffected()
    gtest.Assert(n, 1)
    value, err = db.GetValue("select `passport` from `user` where id=?", 4)
    gtest.Assert(err, nil)
    gtest.Assert(value.String(), "t4")

    // batch with Insert
    if r, err := db.Insert("user", []interface{} {
        map[interface{}]interface{} {
            "id"          : 200,
            "passport"    : "t200",
            "password"    : "25d55ad283aa400af464c76d713c07ad",
            "nickname"    : "T200",
            "create_time" : gtime.Now().String(),
        },
        map[interface{}]interface{} {
            "id"          : 300,
            "passport"    : "t300",
            "password"    : "25d55ad283aa400af464c76d713c07ad",
            "nickname"    : "T300",
            "create_time" : gtime.Now().String(),
        },
    }); err != nil {
        gtest.Fatal(err)
    } else {
        n, _ := r.RowsAffected()
        gtest.Assert(n, 2)
    }

    // clear unnecessary data
    result, err = db.Delete("user", "id>?", 1)
    if err != nil {
        gtest.Fatal(err)
    }
    n, _ = result.RowsAffected()
    gtest.Assert(n, 5)
}

func TestDbBase_BatchInsert(t *testing.T) {
    if r, err := db.BatchInsert("user", g.List {
        {
            "id"          : 2,
            "passport"    : "t2",
            "password"    : "25d55ad283aa400af464c76d713c07ad",
            "nickname"    : "T2",
            "create_time" : gtime.Now().String(),
        },
        {
            "id"          : 3,
            "passport"    : "t3",
            "password"    : "25d55ad283aa400af464c76d713c07ad",
            "nickname"    : "T3",
            "create_time" : gtime.Now().String(),
        },
    }, 1); err != nil {
        gtest.Fatal(err)
    } else {
        n, _ := r.RowsAffected()
        gtest.Assert(n, 2)
    }

    result, err := db.Delete("user", "id>?", 1)
    if err != nil {
        gtest.Fatal(err)
    }
    n, _ := result.RowsAffected()
    gtest.Assert(n, 2)

    // []interface{}
    if r, err := db.BatchInsert("user", []interface{} {
        map[interface{}]interface{} {
            "id"          : 2,
            "passport"    : "t2",
            "password"    : "25d55ad283aa400af464c76d713c07ad",
            "nickname"    : "T2",
            "create_time" : gtime.Now().String(),
        },
        map[interface{}]interface{} {
            "id"          : 3,
            "passport"    : "t3",
            "password"    : "25d55ad283aa400af464c76d713c07ad",
            "nickname"    : "T3",
            "create_time" : gtime.Now().String(),
        },
    }, 1); err != nil {
        gtest.Fatal(err)
    } else {
        n, _ := r.RowsAffected()
        gtest.Assert(n, 2)
    }
}

func TestDbBase_Save(t *testing.T) {
    if _, err := db.Save("user", g.Map{
        "id"          : 1,
        "passport"    : "t1",
        "password"    : "25d55ad283aa400af464c76d713c07ad",
        "nickname"    : "T11",
        "create_time" : gtime.Now().String(),
    }); err != nil {
        gtest.Fatal(err)
    }
}

func TestDbBase_Replace(t *testing.T) {
    if _, err := db.Save("user", g.Map{
        "id"          : 1,
        "passport"    : "t1",
        "password"    : "25d55ad283aa400af464c76d713c07ad",
        "nickname"    : "T111",
        "create_time" : gtime.Now().String(),
    }); err != nil {
        gtest.Fatal(err)
    }
}

func TestDbBase_Update(t *testing.T) {
    if result, err := db.Update("user", "create_time='2010-10-10 00:00:01'", "id=3"); err != nil {
        gtest.Fatal(err)
    } else {
        n, _ := result.RowsAffected()
        gtest.Assert(n, 1)
    }
}

func TestDbBase_GetAll(t *testing.T) {
    if result, err := db.GetAll("SELECT * FROM user WHERE id=?", 1); err != nil {
        gtest.Fatal(err)
    } else {
        gtest.Assert(len(result), 1)
    }
}

func TestDbBase_GetOne(t *testing.T) {
    if record, err := db.GetOne("SELECT * FROM user WHERE passport=?", "t1"); err != nil {
        gtest.Fatal(err)
    } else {
        if record == nil {
            gtest.Fatal("FAIL")
        }
        gtest.Assert(record["nickname"].String(), "T111")
    }
}

func TestDbBase_GetValue(t *testing.T) {
    if value, err := db.GetValue("SELECT id FROM user WHERE passport=?", "t3"); err != nil {
        gtest.Fatal(err)
    } else {
        gtest.Assert(value.Int(), 3)
    }
}

func TestDbBase_GetCount(t *testing.T) {
    if count, err := db.GetCount("SELECT * FROM user"); err != nil {
        gtest.Fatal(err)
    } else {
        gtest.Assert(count, 3)
    }
}

func TestDbBase_GetStruct(t *testing.T) {
    type User struct {
        Id         int
        Passport   string
        Password   string
        NickName   string
        CreateTime gtime.Time
    }
    user := new(User)
    if err := db.GetStruct(user, "SELECT * FROM user WHERE id=?", 3); err != nil {
        gtest.Fatal(err)
    } else {
        gtest.Assert(user.CreateTime.String(), "2010-10-10 00:00:01")
    }
}

func TestDbBase_Delete(t *testing.T) {
    if result, err := db.Delete("user", nil); err != nil {
        gtest.Fatal(err)
    } else {
        n, _ := result.RowsAffected()
        gtest.Assert(n, 3)
    }
}

