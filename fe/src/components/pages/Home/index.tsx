import { MouseEvent } from "react";
import { useAppDispatch } from "../../../app/hooks";
import {
  showCreatePredicateModal,
  showCreateStatementModal,
} from "../../../app/reducers/modal-to-show-slice";

type Props = {};

const Home: React.FC<Props> = (props: Props) => {
  const dispatch = useAppDispatch();

  function handleOnClickCreateStatement(
    event: MouseEvent<HTMLButtonElement>
  ): void {
    dispatch(showCreateStatementModal());
  }

  function handleOnClickCreatePredicate(
    event: MouseEvent<HTMLButtonElement>
  ): void {
    dispatch(showCreatePredicateModal());
  }

  return (
    <>
      <section className="hero is-primary">
        <div className="hero-body">
          <div className="container">
            <h1 className="title is-1">Leibniz</h1>
            <h2 className="subtitle">
              Get started creating your logical universe
            </h2>
          </div>
        </div>
      </section>
      <div className="columns">
        <div className="column">
          <div className="card">
            <div className="card-content">
              <div className="columns">
                <div className="column">
                  <button
                    type="button"
                    className="button is-primary"
                    onClick={handleOnClickCreateStatement}
                  >
                    Create Statement
                  </button>
                </div>
                <div className="column">
                  <button
                    type="button"
                    className="button is-primary"
                    onClick={handleOnClickCreatePredicate}
                  >
                    Create Predicate
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default Home;
