import http from "k6/http";
import { sleep } from "k6";

export let options = {
  stages: [
    { duration: "1m", target: 50 },
    { duration: "3m", target: 200 },
    { duration: "1m", target: 0 },
  ],
};

export default function () {
  http.get("http://10.103.1.51/posts");
  sleep(1);
}
