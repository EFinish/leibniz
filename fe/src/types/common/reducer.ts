// export type SHOWMODALACTIONTYPE = "HIDE" | "SHOW_CREATE_STATEMENT";
export enum ModalToShowType {
  SHOW_CREATE_STATEMENT = "CREATE_STATEMENT",
}

export interface ShowModalReducerAction {
  type: ModalToShowType;
  payload: string | null;
}
