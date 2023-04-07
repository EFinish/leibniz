import React from "react";

type Props = {
  children?: React.ReactNode;
};

const Modal: React.FC<Props> = (props: Props) => {
  return (
      <div className="modal is-active">
        <div className="modal-background"></div>
        <div className="modal-card">
          {props.children}
        </div>
      </div>
  );
};

export default Modal;
