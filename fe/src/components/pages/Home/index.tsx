const Home = () => {
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
            <div className="card-content columns">
              <div className="column">
                <button type="button" className="button is-primary">
                  Create Statement
                </button>
              </div>
              <div className="column">
                <button type="button" className="button is-primary">
                  Create Predicate
                </button>
              </div>
            </div>
          </div>
        </div>
        <div className="column"></div>
        <div className="column"></div>
      </div>
    </>
  );
};

export default Home;
