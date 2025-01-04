import { check } from 'k6';
import http from 'k6/http';

const serviceUrl = 'http://localhost:8082/sort';

export const options = {
    stages: [
        { duration: '30s', target: 100 },
        { duration: '1m', target: 100 },
        { duration: '30s', target: 200 },
        { duration: '1m', target: 200 },
        { duration: '30s', target: 400 },
        { duration: '1m', target: 400 },
        { duration: '30s', target: 800 },
        { duration: '1m', target: 800 },
        { duration: '2m', target: 0 },
    ],
    thresholds: {
        "http_req_failed": ["rate<0.01"],
    },
};

const N = 1000;
const MAX = 1000;

function generateRandomArray(N) {
    return Array.from({ length: N }, () => Math.floor(Math.random() * MAX));
}

function isSorted(arr) {
    for (let i = 1; i < arr.length; i++) {
        if (arr[i - 1] > arr[i]) {
            return false;
        }
    }

    return true;
}

export default function () {
    const randomArray = generateRandomArray(N);
    const payload = JSON.stringify(randomArray);
    const params = { headers: { 'Content-Type': 'application/json' } };

    const res = http.post(serviceUrl, payload, params);

    check(res, {
        'status is 200': (r) => r.status === 200,
        'response is sorted': (r) => isSorted(JSON.parse(r.body)),
    });
}
