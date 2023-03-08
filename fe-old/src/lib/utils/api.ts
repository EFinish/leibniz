import type { CreateSubjectResponse, CreateSubjectRequest, ReadSubjectRequest, ReadSubjectResponse } from "$lib/proto/argumentaccess/v1/server";
import axios from "axios";

export async function CreateSubject(request: CreateSubjectRequest): CreateSubjectResponse {
    axios.post("http://localhost:9000/v1/subject")
}

export async function GetSubjects(request: ReadSubjectRequest): ReadSubjectResponse {
    await axios.get("http://localhost:9000/v1/subject").then(function (response) {
        // handle success
        console.log(response);
      })
      .catch(function (error) {
        // handle error
        console.log(error);
      })
      .finally(function () {
        // always executed
      });

      
}