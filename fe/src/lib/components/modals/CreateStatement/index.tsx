import React from "react";
import { useAppDispatch } from "../../../../app/hooks";
import { hideModal } from "../../../../app/reducers/modal-to-show-slice";
import { CreateSubjectRequest } from "../../../../types/api/request";
import { CreateSubject } from "../../../common/apiWrapper";
import Modal from "../../reusable/Modal";

type Props = {};

const CreateStatementModal: React.FC<Props> = (props: Props) => {
  const dispatch = useAppDispatch();

  const [subjectBody, setSubjectBody] = React.useState("");

  function closeModal(_event: React.SyntheticEvent): void {
    dispatch(hideModal());
  }

  function handleOnChange(event: React.FormEvent<HTMLInputElement>): void {
    setSubjectBody(event.currentTarget.value);
  }

  async function handleOnSubmit(_event: React.SyntheticEvent): Promise<void> {
    const req: CreateSubjectRequest = {
      subject: {
        body: subjectBody,
      },
    };

    const resp = await CreateSubject(req);
    // dispatch(hideModal());
  }

  return (
    <Modal>
      <header className="modal-card-head">
        <p className="modal-card-title">Create Statement</p>
        <button
          className="delete"
          aria-label="close"
          onClick={closeModal}
        ></button>
      </header>
      <section className="modal-card-body">
        <div className="field">
          <label className="label">Subject body</label>
          <div className="control">
            <input
              className="input"
              type="text"
              placeholder="Subject body"
              value={subjectBody}
              onChange={handleOnChange}
            />
            <p className="help">What is the subject?</p>
          </div>
        </div>
      </section>
      <footer className="modal-card-foot">
        <button className="button is-success" onClick={handleOnSubmit}>
          Save changes
        </button>
        <button className="button" onClick={closeModal}>
          Cancel
        </button>
      </footer>
    </Modal>
  );
};

export default CreateStatementModal;
