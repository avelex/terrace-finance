import { Header } from "@/components/Header";
import { VaultsCard } from "@/components/VaultsCard";
import { BridgeCard } from "@/components/BridgeCard";
import { OperationsTable } from "@/components/OperationsTable";

export default function ExplorerPage() {
    return (
        <>
            <Header />
            <main className="flex flex-col gap-8 w-full max-w-[1400px] mx-auto px-[60px] pt-[120px] pb-[60px]">
                <div className="grid grid-cols-2 gap-8 items-start">
                    <VaultsCard />
                    <BridgeCard />
                </div>
                <OperationsTable />
            </main>
        </>
    );
}
