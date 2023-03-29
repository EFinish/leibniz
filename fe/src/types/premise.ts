import { Predicate } from "./predicate";
import { Subject } from "./subject";

export enum SubjectQuantifier {
  ALL = 0,
  SOME = 1,
  NONE = 2,
}
export enum PredicateQualifier {
  NOTISAREHASHAVE = 0,
  ISAREHASHAVE = 1,
}

export type Premise = {
  id?: string;
  subject: Subject;
  subjectQuantifier: SubjectQuantifier;
  predicate: Predicate;
  predicateQualifier: PredicateQualifier;
};
