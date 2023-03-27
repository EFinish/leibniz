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
      state.value = "CREATE_STATEMENT"
    },
    showCreatePredicateModal(state) {
      state.value = "CREATE_PREDICATE"
    }
  },
});

export const { hideModal, showCreateStatementModal, showCreatePredicateModal } = modalToShowSlice.actions;
export default modalToShowSlice.reducer;
