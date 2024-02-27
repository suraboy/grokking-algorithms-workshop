# Load Test Results

## Overview
This document provides an overview of the load test conducted on the binary search and simple search APIs.

## Test Scenario
- **Tool Used**: k6
- **Test Duration**: 10s
- **Number of Virtual Users (VUs)**: 10vus

## Test Results

### Selection Sort API
- **Status**: 200
- **Success Rate**: 100.00% (100 out of 100)
- **Data Received**: 407 kb/s
- **Data Sent**: 405 kB/s
- **HTTP Request Duration**:
  - Average: 1.07s
  - Median: 1.04s
  - Maximum: 1.09s
  - 90th Percentile: 1.08s
  - 95th Percentile: 1.08s
- **HTTP Request Throughput**: 9.308095/s


### Simple Sort API

- **Status**: 200
- **Success Rate**: 100% (90 out of 100)
- **Data Received**: 376 kB/s
- **Data Sent**: 374 kB/s
- **HTTP Request Duration**:
    - Average: 1.15s
    - Median: 1.13s
    - Maximum: 1.17s
    - 90th Percentile: 1.16s
    - 95th Percentile: 1.17s
- **HTTP Request Throughput**: 8.612216/s