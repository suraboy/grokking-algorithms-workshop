import http from 'k6/http';
import { check, group } from 'k6';

export default function () {
    group('Binary Search', function () {
        // Binary Search
        let res = http.get('http://localhost:8080/binary-search?target=11');
        check(res, {
            'status is 200': (r) => r.status === 200,
            'response is correct': (r) => r.body.includes('Binary search result:'),
        });
    });

    // group('Simple Search', function () {
    //     // Simple Search
    //     let res = http.get('http://localhost:8080/simple-search?target=11');
    //     check(res, {
    //         'status is 200': (r) => r.status === 200,
    //         'response is correct': (r) => r.body.includes('Simple search result:'),
    //     });
    // });
}
