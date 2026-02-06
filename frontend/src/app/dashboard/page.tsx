import { Header } from "@/components/Header";
import { DepositsListCard } from "@/components/DepositsListCard";
import { PermitsListCard } from "@/components/PermitsListCard";

export default function DashboardPage() {
    return (
        <>
            <Header />
            <main className="main-container">
                <div className="cards-container">
                    <DepositsListCard />
                    <PermitsListCard />
                </div>
            </main>
        </>
    );
}
