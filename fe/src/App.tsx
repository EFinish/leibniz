import "bulma/css/bulma.min.css";
import { useAppSelector } from "./app/hooks";
import CreatePredicateModal from "./lib/components/modals/CreatePredicate";
import CreateStatementModal from "./lib/components/modals/CreateStatement";
import Home from "./lib/components/pages/Home";

const App = () => {
  const modalToShow = useAppSelector((state) => state.modalToShow.value);

  return (
    <>
      <Home />
      {modalToShow !== null && (
        <div>
          {modalToShow === "CREATE_STATEMENT" && <CreateStatementModal />}
          {modalToShow === "CREATE_PREDICATE" && <CreatePredicateModal />}
        </div>
      )}
    </>
  );
};

export default App;
