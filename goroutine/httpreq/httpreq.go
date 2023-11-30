package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	urls := []string{
		"https://www.google.com/",
		"https://docs.aws.amazon.com/zh_cn/inspector/latest/user/scanning-ec2.html#deep-inspection",
		"https://docs.aws.amazon.com/aws-managed-policy/latest/reference/AmazonSSMManagedInstanceCore.html",
		"https://github.com/Lxb921006/chat",
		"https://stackoverflow.com/",
		"https://docs.aws.amazon.com/systems-manager/latest/userguide/agent-install-rhel.html",
		"https://www.baidu.com/",
		"https://www.qq.com/",
		"https://book.qq.com/book-detail/42859492",
		"https://book.qq.com/book-read/42859492/1",
		"https://book.qq.com/book-read/42859492/2",
		"https://book.qq.com/book-read/42859492/3",
		"https://book.qq.com/book-read/42859492/4",
		"https://book.qq.com/book-read/42859492/5",
		"https://book.qq.com/book-read/42859492/6",
		"https://book.qq.com/book-read/42859492/7",
		"https://book.qq.com/book-read/42859492/8",
		"https://book.qq.com/book-read/42859492/9",
		"https://book.qq.com/book-read/42859492/10",
		"https://book.qq.com/book-read/42859492/11",
		"https://book.qq.com/book-read/42859492/12",
		"https://book.qq.com/book-read/42859492/13",
		"https://book.qq.com/book-read/42859492/14",
		"https://book.qq.com/book-read/42859492/15",
		"https://book.qq.com/book-read/42859492/16",
		"https://book.qq.com/book-read/42859492/17",
		"https://book.qq.com/book-read/42859492/18",
		"https://book.qq.com/book-read/42859492/19",
		"https://book.qq.com/book-read/42859492/20",
		"https://book.qq.com/book-read/42859492/21",
		"https://book.qq.com/book-read/42859492/22",
		"https://book.qq.com/book-read/42859492/23",
		"https://book.qq.com/book-read/42859492/24",
		"https://book.qq.com/book-read/42859492/25",
		"https://book.qq.com/book-read/42859492/26",
		"https://book.qq.com/book-read/42859492/27",
		"https://book.qq.com/book-read/42859492/28",
		"https://book.qq.com/book-read/42859492/29",
		"https://book.qq.com/book-read/42859492/30",
		"https://book.qq.com/book-read/42859492/31",
		"https://book.qq.com/book-read/42859492/32",
		"https://book.qq.com/book-read/42859492/33",
		"https://book.qq.com/book-read/42859492/34",
		"https://book.qq.com/book-read/42859492/35",
		"https://book.qq.com/book-read/42859492/36",
		"https://book.qq.com/book-read/42859492/37",
		"https://book.qq.com/book-read/42859492/38",
		"https://book.qq.com/book-read/42859492/39",
		"https://book.qq.com/book-read/42859492/40",
		"https://book.qq.com/book-read/42859492/41",
		"https://book.qq.com/book-read/42859492/42",
		"https://book.qq.com/book-read/42859492/43",
		"https://book.qq.com/book-read/42859492/44",
		"https://book.qq.com/book-read/42859492/45",
		"https://book.qq.com/book-read/42859492/46",
		"https://book.qq.com/book-read/42859492/47",
		"https://book.qq.com/book-read/42859492/48",
		"https://book.qq.com/book-read/42859492/49",
		"https://book.qq.com/book-read/42859492/50",
		"https://book.qq.com/book-read/42859492/51",
		"https://book.qq.com/book-read/42859492/52",
		"https://book.qq.com/book-read/42859492/53",
		"https://book.qq.com/book-read/42859492/54",
		"https://book.qq.com/book-read/42859492/55",
		"https://book.qq.com/book-read/42859492/56",
		"https://book.qq.com/book-read/42859492/57",
		"https://book.qq.com/book-read/42859492/58",
		"https://book.qq.com/book-read/42859492/59",
		"https://book.qq.com/book-read/42859492/60",
		"https://book.qq.com/book-read/42859492/61",
		"https://book.qq.com/book-read/42859492/62",
		"https://book.qq.com/book-read/42859492/63",
		"https://book.qq.com/book-read/42859492/64",
		"https://book.qq.com/book-read/42859492/65",
		"https://book.qq.com/book-read/42859492/66",
		"https://book.qq.com/book-read/42859492/67",
		"https://book.qq.com/book-read/42859492/68",
		"https://book.qq.com/book-read/42859492/69",
		"https://book.qq.com/book-read/42859492/70",
		"https://book.qq.com/book-read/42859492/71",
		"https://book.qq.com/book-read/42859492/72",
		"https://book.qq.com/book-read/42859492/73",
		"https://book.qq.com/book-read/42859492/74",
		"https://book.qq.com/book-read/42859492/75",
		"https://book.qq.com/book-read/42859492/76",
		"https://book.qq.com/book-read/42859492/77",
		"https://book.qq.com/book-read/42859492/78",
		"https://book.qq.com/book-read/42859492/79",
		"https://book.qq.com/book-read/42859492/80",
		"https://book.qq.com/book-read/42859492/81",
		"https://book.qq.com/book-read/42859492/82",
		"https://book.qq.com/book-read/42859492/83",
		"https://book.qq.com/book-read/42859492/84",
		"https://book.qq.com/book-read/42859492/85",
		"https://book.qq.com/book-read/42859492/86",
		"https://book.qq.com/book-read/42859492/87",
		"https://book.qq.com/book-read/42859492/88",
		"https://book.qq.com/book-read/42859492/89",
		"https://book.qq.com/book-read/42859492/90",
		"https://book.qq.com/book-read/42859492/91",
		"https://book.qq.com/book-read/42859492/92",
		"https://book.qq.com/book-read/42859492/93",
		"https://book.qq.com/book-read/42859492/94",
		"https://book.qq.com/book-read/42859492/95",
		"https://book.qq.com/book-read/42859492/96",
		"https://book.qq.com/book-read/42859492/97",
		"https://book.qq.com/book-read/42859492/98",
		"https://book.qq.com/book-read/42859492/99",
		"https://book.qq.com/book-read/42859492/100",
		"https://book.qq.com/book-read/42859492/101",
		"https://book.qq.com/book-read/42859492/102",
		"https://book.qq.com/book-read/42859492/103",
		"https://book.qq.com/book-read/42859492/104",
		"https://book.qq.com/book-read/42859492/105",
		"https://book.qq.com/book-read/42859492/106",
		"https://book.qq.com/book-read/42859492/107",
		"https://book.qq.com/book-read/42859492/108",
		"https://book.qq.com/book-read/42859492/109",
		"https://book.qq.com/book-read/42859492/110",
		"https://book.qq.com/book-read/42859492/111",
		"https://book.qq.com/book-read/42859492/112",
		"https://book.qq.com/book-read/42859492/113",
		"https://book.qq.com/book-read/42859492/114",
		"https://book.qq.com/book-read/42859492/115",
		"https://book.qq.com/book-read/42859492/116",
		"https://book.qq.com/book-read/42859492/117",
		"https://book.qq.com/book-read/42859492/118",
		"https://book.qq.com/book-read/42859492/119",
		"https://book.qq.com/book-read/42859492/120",
		"https://book.qq.com/book-read/42859492/121",
		"https://book.qq.com/book-read/42859492/122",
		"https://book.qq.com/book-read/42859492/123",
		"https://book.qq.com/book-read/42859492/124",
		"https://book.qq.com/book-read/42859492/125",
		"https://book.qq.com/book-read/42859492/126",
		"https://book.qq.com/book-read/42859492/127",
		"https://book.qq.com/book-read/42859492/128",
		"https://book.qq.com/book-read/42859492/129",
		"https://book.qq.com/book-read/42859492/130",
		"https://book.qq.com/book-read/42859492/131",
		"https://book.qq.com/book-read/42859492/132",
		"https://book.qq.com/book-read/42859492/133",
		"https://book.qq.com/book-read/42859492/134",
		"https://book.qq.com/book-read/42859492/135",
		"https://book.qq.com/book-read/42859492/136",
		"https://book.qq.com/book-read/42859492/137",
		"https://book.qq.com/book-read/42859492/138",
		"https://book.qq.com/book-read/42859492/139",
		"https://book.qq.com/book-read/42859492/140",
		"https://book.qq.com/book-read/42859492/141",
		"https://book.qq.com/book-read/42859492/142",
		"https://book.qq.com/book-read/42859492/143",
		"https://book.qq.com/book-read/42859492/144",
		"https://book.qq.com/book-read/42859492/145",
		"https://book.qq.com/book-read/42859492/146",
		"https://book.qq.com/book-read/42859492/147",
		"https://book.qq.com/book-read/42859492/148",
		"https://book.qq.com/book-read/42859492/149",
		"https://book.qq.com/book-read/42859492/150",
		"https://book.qq.com/book-read/42859492/151",
		"https://book.qq.com/book-read/42859492/152",
		"https://book.qq.com/book-read/42859492/153",
		"https://book.qq.com/book-read/42859492/154",
		"https://book.qq.com/book-read/42859492/155",
		"https://book.qq.com/book-read/42859492/156",
		"https://book.qq.com/book-read/42859492/157",
		"https://book.qq.com/book-read/42859492/158",
		"https://book.qq.com/book-read/42859492/159",
		"https://book.qq.com/book-read/42859492/160",
		"https://book.qq.com/book-read/42859492/161",
		"https://book.qq.com/book-read/42859492/162",
		"https://book.qq.com/book-read/42859492/163",
		"https://book.qq.com/book-read/42859492/164",
		"https://book.qq.com/book-read/42859492/165",
		"https://book.qq.com/book-read/42859492/166",
		"https://book.qq.com/book-read/42859492/167",
		"https://book.qq.com/book-read/42859492/168",
		"https://book.qq.com/book-read/42859492/169",
		"https://book.qq.com/book-read/42859492/170",
		"https://book.qq.com/book-read/42859492/171",
		"https://book.qq.com/book-read/42859492/172",
		"https://book.qq.com/book-read/42859492/173",
		"https://book.qq.com/book-read/42859492/174",
		"https://book.qq.com/book-read/42859492/175",
		"https://book.qq.com/book-read/42859492/176",
		"https://book.qq.com/book-read/42859492/177",
		"https://book.qq.com/book-read/42859492/178",
		"https://book.qq.com/book-read/42859492/179",
		"https://book.qq.com/book-read/42859492/180",
		"https://book.qq.com/book-read/42859492/181",
		"https://book.qq.com/book-read/42859492/182",
		"https://book.qq.com/book-read/42859492/183",
		"https://book.qq.com/book-read/42859492/184",
		"https://book.qq.com/book-read/42859492/185",
		"https://book.qq.com/book-read/42859492/186",
		"https://book.qq.com/book-read/42859492/187",
		"https://book.qq.com/book-read/42859492/188",
		"https://book.qq.com/book-read/42859492/189",
		"https://book.qq.com/book-read/42859492/190",
		"https://book.qq.com/book-read/42859492/191",
		"https://book.qq.com/book-read/42859492/192",
		"https://book.qq.com/book-read/42859492/193",
		"https://book.qq.com/book-read/42859492/194",
		"https://book.qq.com/book-read/42859492/195",
		"https://book.qq.com/book-read/42859492/196",
		"https://book.qq.com/book-read/42859492/197",
		"https://book.qq.com/book-read/42859492/198",
		"https://book.qq.com/book-read/42859492/199",
		"https://book.qq.com/book-read/42859492/200",
		"https://book.qq.com/book-read/42859492/201",
		"https://book.qq.com/book-read/42859492/202",
		"https://book.qq.com/book-read/42859492/203",
		"https://book.qq.com/book-read/42859492/204",
		"https://book.qq.com/book-read/42859492/205",
		"https://book.qq.com/book-read/42859492/206",
		"https://book.qq.com/book-read/42859492/207",
		"https://book.qq.com/book-read/42859492/208",
		"https://book.qq.com/book-read/42859492/209",
		"https://book.qq.com/book-read/42859492/210",
		"https://book.qq.com/book-read/42859492/211",
		"https://book.qq.com/book-read/42859492/212",
		"https://book.qq.com/book-read/42859492/213",
		"https://book.qq.com/book-read/42859492/214",
		"https://book.qq.com/book-read/42859492/215",
		"https://book.qq.com/book-read/42859492/216",
		"https://book.qq.com/book-read/42859492/217",
		"https://book.qq.com/book-read/42859492/218",
		"https://book.qq.com/book-read/42859492/219",
		"https://book.qq.com/book-read/42859492/220",
		"https://book.qq.com/book-read/42859492/221",
		"https://book.qq.com/book-read/42859492/222",
		"https://book.qq.com/book-read/42859492/223",
		"https://book.qq.com/book-read/42859492/224",
		"https://book.qq.com/book-read/42859492/225",
		"https://book.qq.com/book-read/42859492/226",
		"https://book.qq.com/book-read/42859492/227",
		"https://book.qq.com/book-read/42859492/228",
		"https://book.qq.com/book-read/42859492/229",
		"https://book.qq.com/book-read/42859492/230",
		"https://book.qq.com/book-read/42859492/231",
		"https://book.qq.com/book-read/42859492/232",
		"https://book.qq.com/book-read/42859492/233",
		"https://book.qq.com/book-read/42859492/234",
		"https://book.qq.com/book-read/42859492/235",
		"https://book.qq.com/book-read/42859492/236",
		"https://book.qq.com/book-read/42859492/237",
		"https://book.qq.com/book-read/42859492/238",
		"https://book.qq.com/book-read/42859492/239",
		"https://book.qq.com/book-read/42859492/240",
		"https://book.qq.com/book-read/42859492/241",
		"https://book.qq.com/book-read/42859492/242",
		"https://book.qq.com/book-read/42859492/243",
		"https://book.qq.com/book-read/42859492/244",
		"https://book.qq.com/book-read/42859492/245",
		"https://book.qq.com/book-read/42859492/246",
		"https://book.qq.com/book-read/42859492/247",
		"https://book.qq.com/book-read/42859492/248",
		"https://book.qq.com/book-read/42859492/249",
		"https://book.qq.com/book-read/42859492/250",
		"https://book.qq.com/book-read/42859492/251",
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var workers = 10
	var countFinishedWorker int

	req := newRequest(ctx, workers)
	req.wg.Add(1)

	req.start()

	go func() {
		for {
			select {
			case <-req.ctx.Done():
				return
			case <-req.finished:
				countFinishedWorker++
				if countFinishedWorker == workers {
					close(req.result)
					return
				}

			}
		}
	}()

	go func() {
		defer req.wg.Done()
		for {
			select {
			case <-req.ctx.Done():
				return
			case v, ok := <-req.result:
				if !ok {
					return
				}
				fmt.Println(v)
			}
		}
	}()

	go func() {
		for _, v := range urls {
			req.task <- v
		}
		req.stop()
	}()

	req.wg.Wait()
}

type result struct {
	url  string
	resp string
}

//type task func() *result

type request struct {
	ctx      context.Context
	task     chan string
	result   chan *result
	workers  int
	finished chan struct{}
	once     *sync.Once
	wg       *sync.WaitGroup
	lock     *sync.Mutex
}

func newRequest(ctx context.Context, workers int) *request {
	return &request{
		ctx:      ctx,
		task:     make(chan string),
		result:   make(chan *result),
		finished: make(chan struct{}),
		once:     new(sync.Once),
		wg:       new(sync.WaitGroup),
		lock:     new(sync.Mutex),
		workers:  workers,
	}
}

func (r *request) start() {
	for i := 0; i < r.workers; i++ {
		go r.work()
	}
}

func (r *request) run(u string) *result {
	r.lock.Lock()
	defer r.lock.Unlock()

	ctx1, cancel1 := context.WithTimeout(context.Background(), time.Second)
	defer cancel1()

	var req = new(result)
	var done = make(chan *http.Response)

	go func() {
		resp, _ := http.Get(u)
		done <- resp
	}()

	select {
	case <-ctx1.Done():
		req.url = u
		req.resp = "time out"
	case v := <-done:
		req.url = v.Request.URL.String()
		req.resp = v.Request.Host
	}

	return req
}

func (r *request) stop() {
	close(r.task)
}

func (r *request) work() {
	for {
		select {
		case <-r.ctx.Done():
			return
		case v, ok := <-r.task:
			if !ok {
				r.finished <- struct{}{}
				return
			}

			resp := r.run(v)
			r.result <- resp

		}
	}
}
