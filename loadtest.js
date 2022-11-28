import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
  stages: [
    { duration: '1m', target: 500 }, // simulate ramp-up of traffic from 1 to 500 users over 1 minutes.
    { duration: '5m', target: 1000 }, // stay at 1000 users for 5 minutes
  ],
};

const ports = ['14568', '14569'];

export default function () {
  const port = ports[getRandomInt(0,2)]  
  const url = 'http://127.0.0.1:'+port+'/api/v1/study-asynq/announcement/schedule';
  const payload = JSON.stringify({
    from: getRandomStr(),
    message: getRandomStr(),
    seconds: getRandomInt(5,10)
  });
  const params = {
    headers: {
    'Content-Type': 'application/json',
    },
  };
  
  http.post(url, payload, params);
  sleep(1);
}

// a max number is not included
function getRandomInt(min, max) {
  min = Math.ceil(min);
  max = Math.floor(max);
  return Math.floor(Math.random() * (max - min)) + min;
}

function getRandomStr() {
    return Math.random().toString(36).substring(2, 12);
}