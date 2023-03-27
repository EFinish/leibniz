import "bulma/css/bulma.min.css";
import { useAppSelector } from "./app/hooks";
import CreateStatementModal from "./components/modals/CreateStatement";
import Home from "./components/pages/Home";
import { ModalToShowType, ShowModalReducerAction } from "./types/common/reducer";

const App = () => {
  const modalToShow = useAppSelector((state) => state.modalToShow.value);

  return (
    <>
      <Home />
        {modalToShow !== null && Object.values(ModalToShowType).includes(modalToShow) && (
          <CreateStatementModal />
        )}
    </>
  );
};

export default App;
