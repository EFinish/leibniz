import { createSlice } from '@reduxjs/toolkit';
import { ModalToShowType } from '../../types/common/reducer';

interface ModalToShowState {
  value: ModalToShowType | null;
}

const initialState: ModalToShowState = {
  value: null,
};

const modalToShowSlice = createSlice({
  name: 'modalToShow',
  initialState,
  reducers: {
    hideModal(state) {
        state.value = null
    },
    showCreateStatementModal(state) {
      state.value = ModalToShowType.SHOW_CREATE_STATEMENT
    }
  },
});

export const { hideModal, showCreateStatementModal } = modalToShowSlice.actions;
export default modalToShowSlice.reducer;
