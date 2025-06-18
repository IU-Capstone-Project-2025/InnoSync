import Navbar from "@/components/Navbar";
import ProfileCreationForm from "@/components/ProfileCreationForm";

const Index = () => {
  return (
    <div className="min-h-screen w-full bg-white font-montserrat">
      {/* Navbar */}
      <Navbar />

      {/* Main Content */}
      <main className="flex flex-col items-center pt-8 sm:pt-16 lg:pt-[190px] pb-8 sm:pb-16 lg:pb-[290px] px-4">
        <ProfileCreationForm />
      </main>
    </div>
  );
};

export default Index;
