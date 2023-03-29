import { Premise } from "./premise";
import { Proposition } from "./proposition";

export type ConditionalStatement = {
  id?: string;
  ifPremise: Premise;
  thenPremise: Premise;
  ifProposition: Proposition;
  thenProposition: Proposition;
};
