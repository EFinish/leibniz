import React, { useEffect, useState } from "react";
import { createPortal } from 'react-dom';

type Props = {
  children?: React.ReactNode;
};

const modalRootId = "react-portal-wrapper"
const modalRoot = document.getElementById(modalRootId) as HTMLElement;

const ReactPortal: React.FC<Props> = ({children}) => {
    const wrapperElement = document.createElement('div');

    useEffect(() => {
        modalRoot.appendChild(wrapperElement);
        return () => {
            modalRoot.removeChild(wrapperElement);
        }
    }, [])

    if (wrapperElement === null) return null;

    return createPortal(children, wrapperElement);

}

export default ReactPortal;
