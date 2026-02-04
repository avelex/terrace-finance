import { Header } from "@/components/Header";
import { DepositCard } from "@/components/DepositCard";
import { BalancesCard } from "@/components/BalancesCard";

export default function Home() {
  return (
    <>
      <Header />
      <main className="main-container">
        <div className="cards-container">
          <BalancesCard />
          <DepositCard />
        </div>
      </main>
    </>
  );
}
