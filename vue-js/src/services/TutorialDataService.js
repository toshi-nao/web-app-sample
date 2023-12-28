import http from "../http-common";

class TutorialDataService {
  getAll() {
    return http.get("/tutorials");
  }

  get(Id) {
    return http.get(`/tutorials/${Id}`);
  }

  create(data) {
    return http.post("/tutorials", data);
  }

  update(Id, data) {
    return http.put(`/tutorials/${Id}`, data);
  }

  delete(Id) {
    return http.delete(`/tutorials/${Id}`);
  }

  deleteAll() {
    return http.delete(`/tutorials`);
  }

  findByTitle(title) {
    return http.get(`/tutorials?title=${title}`);
  }
}

export default new TutorialDataService();