import { Header } from "@/components/Header";
import { DepositCard } from "@/components/DepositCard";

export default function Home() {
  return (
    <>
      <Header />
      <main className="main-container">
        <div className="content-wrapper">
          <div className="hero-section">
            <h1 className="hero-title">Terrace</h1>
            <p className="hero-subtitle">DECENTRALIZED ONE-CLICK STABLECOIN YIELD PLATFORM</p>
          </div>
          <div className="card-section">
            <DepositCard />
          </div>

        </div>
      </main>
    </>
  );
}

