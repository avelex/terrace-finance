import { Header } from "@/components/Header";
import { VaultsCard } from "@/components/VaultsCard";
import { BridgeCard } from "@/components/BridgeCard";
import { OperationsTable } from "@/components/OperationsTable";

export default function ExplorerPage() {
    return (
        <>
            <Header />
            <main className="main-container explorer-container">
                <div className="explorer-top-row">
                    <VaultsCard />
                    <BridgeCard />
                </div>
                <OperationsTable />
            </main>
        </>
    );
}
