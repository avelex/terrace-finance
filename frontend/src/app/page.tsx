import { Header } from "@/components/Header";
import { DepositCard } from "@/components/DepositCard";
import { BalancesCard } from "@/components/BalancesCard";
import { DepositsListCard } from "@/components/DepositsListCard";
import { PermitsListCard } from "@/components/PermitsListCard";

export default function Home() {
  return (
    <>
      <Header />
      <main className="main-container">
        <div className="cards-container">
          <BalancesCard />
          <DepositsListCard />
          <PermitsListCard />
          <DepositCard />
        </div>
      </main>
    </>
  );
}
