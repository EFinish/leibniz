import { ConditionalStatement } from "./conditional_statement";
import { Premise } from "./premise";

export type Argument = {
  id?: string;
  title: string;
  premises: Premise[];
  conditionalStatements: ConditionalStatement[];
  conclusionPremise: Premise;
};
