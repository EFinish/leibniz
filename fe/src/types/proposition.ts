import { Premise } from "./premise";

export enum PropositionType {
  AND = 0,
  OR = 1,
  NAND = 2,
  NOR = 3,
  XOR = 4,
  XNOR = 5,
}

export type Proposition = {
  id?: string;
  propositionType: PropositionType;
  propositionSubPremises: Premise[];
  propositionSubPropositions: Proposition[];
};
