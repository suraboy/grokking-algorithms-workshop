# Load Test Results

## Overview
This document provides an overview of the load test conducted on the binary search and simple search APIs.

## Test Scenario
- **Tool Used**: k6
- **Test Duration**: 5s
- **Number of Virtual Users (VUs)**: 2vus

## Test Results

### Binary Search API
- **Status**: 200
- **Success Rate**: 100.00% (204790 out of 204790)
- **Data Received**: 15 MB (2.9 MB/s)
- **Data Sent**: 11 MB (2.1 MB/s)
- **HTTP Request Duration**:
  - Average: 54.61µs
  - Median: 49µs
  - Maximum: 71.69ms
  - 90th Percentile: 67µs
  - 95th Percentile: 79µs
- **HTTP Request Throughput**: 20472.407885/s


### Simple Search API

- **Status**: 200
- **Success Rate**: 100% (195958 out of 195958)
- **Data Received**: 14 kB
- **Data Sent**: 10 kB
- **HTTP Request Duration**:
    - Average: 90.43µs
    - Median: 79.95µs
    - Maximum: 8.44ms
    - 90th Percentile: 110.66µs
    - 95th Percentile: 133.17µs
- **HTTP Request Throughput**: 19589.104444/s