# **Introduction**

This repository demonstrates the performance difference and advantages of using Go with MongoDB, and how it can be incorporated into our development stack. The HTTP server is bootstrapped using [Gin](https://gin-gonic.com/)
framework.

Gin by default executes each incoming request in a separate go routine (lightweight thread). This allows the framework to handle high levels of concurrency and process multiple requests simultaneously without blocking.

---

---

## **Setup**

- Install Node.js and npm.
- Install the latest version of Go.
- Make sure you have axisPayCore database and transactionDetails collection in your local mongo server
- Run both services. You can start the Go server by navigating to the project directory and typing "go run main.go".

---

# **K6 Test Results**

I used k6, a load testing tool built by the Grafana team, for the load testing experiment. This tool is built with Go and does not operate inside the Node runtime, which is why it supports concurrent requests with its virtual user concept.

> **Note**: all the tests ran for 6 mins with 20 concurrent virtual user

##

##

### case 1 → Fetching 20 records per request.

- node

# ![node](./media/v2/20_node.png)

- go

# ![node](./media/v2/20_go.png)

the difference here is noticeable in the response time but not that significant because the operation is not cpu intensive memory consuming and the response payload is not big

### case 2 → Fetching 100 records per request.

- node

# ![node](./media/v2/100_node_a.png)

- go

# ![node](./media/v2/100_go.png)

the same happens here node still handles the load well, but there's an increase in the response time, go still wins

### case 3 → Fetching 1000 records per request.

- node

# ![node](./media/v2/1000_node_a.png)

# ![node](./media/v2/1000_node_b.png)

- go

# ![node](./media/v2/1000_go.png)

Go in this example handles the load very well with no issues, node on the other hand struggles with big payloads queries due to its single-threaded nature

##

##

# Postman Test Results (_for the sake of it_)

### first, fetching a single document with both servers

## ![node](./media//node_single.gif)

---

---

## ![go](./media//go_single.gif)

---

---

### fetching 1000 document

![go-node](./media//go-node%20-1000.gif)

---

---

### fetching 10,000 documents in node

## ![node](./media//node_10_000.gif)

---

---

### fetching 10,000 documents in go

## ![node](./media//go_10_000.gif)

---

---

---

---

## Load Testing With Prometheus and Grafana

---

In order to measure the amount of memory and cpu time consumed by every service I created this project using the following technologies:

- **Prometheus**: for querying metrics info from each service by exposing a "/metrics" API that's called by prometheus server to pull cpu and memory data.
- **Grafana**: for displaying the data gathered by prometheus in a clear dashboard
- **Apache Bench**: for executing the load testing scripts locally

---

### Test Results

<!--
| Language | Color                                    |
| -------- | ---------------------------------------- |
| Node     | <span style="color:blue">blue</span>     |
| Java     | <span style="color:green">green</span>   |
| Go       | <span style="color:yellow">yellow</span> | -->

<!-- --- -->

- Fetching 3000 requests wih 20 concurrent connections in node
  ![](./media/grafana_tests/node_graph.png)
  ![](./media/grafana_tests/node_nums.png)

  as shown here node is not the best option for this kind of operation because it is a single threaded language and it's not that fast in terms of execution time and for the cpu and memory usage it is not that good either.

---

- Fetching 3000 requests wih 20 concurrent connections in Go
  ![](./media/grafana_tests/go_graph.png)

  > _zoomed-in for better visibility_

  ![](./media/grafana_tests/go_graph_2.png)
  ![](./media/grafana_tests/go_nums.png)

  go here runs smoothly in all aspects, it's fast, it's memory efficient and it's cpu efficient.

---

- Fetching 3000 requests wih 20 concurrent connections in Java
  ![](./media/grafana_tests/java_graph.png)
  ![](./media/grafana_tests/java_nums.png)

  java is good in terms of execution time but it's not that good in terms of memory consumption and cpu usage compared to go

---

- Fetching 5000 requests wih 20 concurrent connections all services
  ![](./media/grafana_tests/go_node_java_5000.png)

When I ran this test, Node and Go executed as expected. However, Node's response time and memory footprint increased, while Go's remained the same. Java, on the other hand, experienced a huge spike in memory usage even after the test was completed, due to its garbage collector.

Both Java and Go utilize garbage collection to manage memory. However, Go's garbage collector is designed to be more efficient and have less impact on CPU and memory usage compared to Java's.
Go's concurrent garbage collector reduces the impact of garbage collection on resources and allows the application to run smoothly without pausing.

The Java garbage collector identifies objects that are no longer being used by the program and frees the memory associated with those objects. However, to do this, the garbage collector needs to temporarily allocate additional memory to perform its analysis and bookkeeping.
