import http from 'k6/http';
import {check, sleep} from 'k6';

export let options = {
    vus: 10, // Number of virtual users
    duration: '10s', // Test duration
};

export default function () {
    // Test selection sort with random arrays of different sizes
    testSort('/selection-sort', generateRandomArray(100), 'Sorted array (Selection Sort):');
    testSort('/selection-sort', generateRandomArray(1000), 'Sorted array (Selection Sort):');
    testSort('/selection-sort', generateRandomArray(10000), 'Sorted array (Selection Sort):');

    // Test simple sort with random arrays of different sizes
    testSort('/simple-sort', generateRandomArray(100), 'Sorted array (Simple Sort):');
    testSort('/simple-sort', generateRandomArray(1000), 'Sorted array (Simple Sort):');
    testSort('/simple-sort', generateRandomArray(10000), 'Sorted array (Simple Sort):');

    // Add more test cases as needed

    // Sleep for a short duration between iterations
    sleep(1);
}

function generateRandomArray(size) {
    let arr = [];
    for (let i = 0; i < size; i++) {
        arr.push(Math.floor(Math.random() * 1000));
    }
    return arr;
}

function testSort(endpoint, arr,msg) {
    let res = http.get(`http://localhost:8080${endpoint}?arr=${arr.join(',')}`);
    check(res, {
        'status is 200': (r) => r.status === 200,
        'response is correct': (r) => r.body.includes(msg),
    });
}
