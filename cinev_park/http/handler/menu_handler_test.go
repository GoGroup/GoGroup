package handler

// func TestBook(t *testing.T) {

// 	tmpl := template.Must(template.ParseGlob("../../../view/template/*"))

// 	schdlRepo := schrep.NewMockScheduleRepo(nil)
// 	schdlServ := schser.NewScheduleService(schdlRepo)

// 	bookRepo := brep.NewMockBookingepo(nil)
// 	bookSer := bser.NewBookingService(bookRepo)

// 	hrep := hallrepo.NewMockHallRepo(nil)
// 	hser := hallser.NewHallService(hrep)

// 	cinr := cinrep.NewMockCinemaRepo(nil)
// 	cins := cinser.NewCinemaService(cinr)

// 	userRepo := urep.NewMockUserRepo(nil)
// 	userServ := user.NewUserService(userRepo)

// 	menuBookHandler := MenuHandler(tmpl, cins, hser, schdlServ, nil, userServ, nil)

// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/theater/schedule/book/", menuBookHandler.TheaterScheduleBook)
// 	ts := httptest.NewTLSServer(mux)
// 	defer ts.Close()

// 	tc := ts.Client()
// 	URL := ts.URL
// 	form := url.Values{}
// 	form.Add("seat", string(model.HallMock.Price))

// 	resp, err := tc.PostForm(sURL+"/theater/schedule/book/Mock Cinema1/1/1", form)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if resp.StatusCode != http.StatusOK {
// 		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
// 	}

// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

//}
