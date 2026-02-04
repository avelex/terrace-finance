import { Header } from "@/components/Header";
import { DepositCard } from "@/components/DepositCard";

export default function Home() {
  return (
    <>
      <Header />
      <main className="main-container">
        <DepositCard />
      </main>
    </>
  );
}
