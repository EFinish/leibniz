import React, { MouseEvent } from "react";
import { useAppDispatch } from "../../../app/hooks";
import { hideModal } from "../../../app/reducers/modal-to-show-slice";
import Modal from "../../common/Modal";

type Props = {};

const CreatePredicateModal: React.FC<Props> = (props: Props) => {
  const dispatch = useAppDispatch();

  function closeModal(
    event: MouseEvent<HTMLButtonElement>
  ): void {
    dispatch(hideModal());
  }

  return (
    <Modal>
      <header className="modal-card-head">
        <p className="modal-card-title">Create Predicate</p>
        <button className="delete" aria-label="close" onClick={closeModal}></button>
      </header>
      <section className="modal-card-body">Create Predicate</section>
      <footer className="modal-card-foot">
        <button className="button is-success">Save changes</button>
        <button className="button" onClick={closeModal}>Cancel</button>
      </footer>
    </Modal>
  );
};

export default CreatePredicateModal;
