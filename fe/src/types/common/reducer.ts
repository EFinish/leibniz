export type ModalToShowType = "CREATE_STATEMENT" | "CREATE_PREDICATE";

export interface ShowModalReducerAction {
  type: ModalToShowType;
  payload: string | null;
}
