import { configureStore } from '@reduxjs/toolkit';
import modalToShowReducer from './reducers/modal-to-show-slice';

export const store = configureStore({
  reducer: {
    modalToShow: modalToShowReducer,
  }
});

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
