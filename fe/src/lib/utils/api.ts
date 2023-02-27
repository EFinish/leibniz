import { CreateSubjectResponse, type CreateSubjectRequest } from "$lib/proto/argumentaccess/v1/server";
import axios from "axios";

export function CreateSubject(request: CreateSubjectRequest): CreateSubjectResponse {
    axios.post("localhost:9002/")
}