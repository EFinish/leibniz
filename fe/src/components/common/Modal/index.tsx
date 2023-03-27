import React from "react";
import ReactPortal from "../../Portal";

type Props = {
  children?: React.ReactNode;
};

const Modal: React.FC<Props> = (props: Props) => {
  return (
    <ReactPortal>
      <div className="modal is-active">
        <div className="modal-background"></div>
        <div className="modal-card">
          {props.children}
        </div>
      </div>
    </ReactPortal>
  );
};

export default Modal;
