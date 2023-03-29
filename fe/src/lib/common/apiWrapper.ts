import { CreateSubjectRequest } from "../../types/api/request";
import { CreateSubjectResponse } from "../../types/api/response";
import { apiPost } from "./api";

export function CreateSubject(
  req: CreateSubjectRequest
): Promise<CreateSubjectResponse> {
  return apiPost("http://localhost:9102/v1/subject", req).then(
    async (resp: Response) => {
      if (!resp.ok) {
        throw new Error(resp.statusText);
      }

      return await (resp.json() as Promise<CreateSubjectResponse>);
    }
  );
}
