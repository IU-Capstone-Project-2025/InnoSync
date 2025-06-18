import { useState } from "react";
import { Plus, ArrowRight } from "lucide-react";

const ProfileCreationForm = () => {
  const [formData, setFormData] = useState({
    fullName: "Ahmed Baha Eddine Alimi",
    email: "3llimi69@gmail.com",
    telegram: "Alias",
    github: "Alias",
    bio: "Bio",
  });

  const handleInputChange = (field: string, value: string) => {
    setFormData((prev) => ({ ...prev, [field]: value }));
  };

  return (
    <div className="flex w-full max-w-[1000px] min-h-[500px] flex-col lg:flex-row justify-center items-start rounded-[20px] bg-[#E4E4E4] shadow-[0px_4px_4px_4px_rgba(0,0,0,0.25)] mx-4">
      {/* Left Side - Profile Picture */}
      <div className="flex w-full lg:w-[400px] min-h-[300px] lg:h-[500px] p-8 lg:p-[103px_96px] flex-col justify-center items-center flex-shrink-0 rounded-t-[30px] lg:rounded-[30px_0px_0px_30px] bg-white">
        <div className="flex w-[208px] flex-col items-center gap-[37px]">
          {/* Profile Picture Upload Area */}
          <div className="flex h-[200px] p-[61px_81px] flex-col justify-center items-center gap-[10px] w-full rounded-full border-4 border-dashed border-inno-lightgreen bg-[rgba(230,230,230,0.42)] cursor-pointer hover:bg-[rgba(230,230,230,0.6)] transition-colors">
            <Plus
              className="w-16 h-16 text-inno-green font-light"
              strokeWidth={1}
            />
          </div>

          {/* User Info */}
          <div className="flex flex-col items-center gap-[9px] w-full">
            <div className="text-black text-center font-montserrat text-xl font-light">
              A.Baha Alimi
            </div>
            <div className="text-black text-center font-montserrat text-xl font-light">
              3llimi69@gmail.com
            </div>
          </div>
        </div>
      </div>

      {/* Right Side - Form */}
      <div className="w-full lg:w-[598px] min-h-[500px] flex-shrink-0 rounded-b-[30px] lg:rounded-[0px_30px_30px_0px] bg-white relative">
        <div className="flex w-full max-w-[475px] flex-col items-start gap-4 absolute left-4 lg:left-[61px] top-[27px] right-4 lg:right-auto px-4 lg:px-0">
          {/* Full Name Field */}
          <div className="flex flex-col justify-center items-center gap-[5px] w-full">
            <label className="w-full text-black font-montserrat text-lg sm:text-xl font-normal">
              Full Name:
            </label>
            <div className="flex h-10 px-6 py-2 items-center gap-[10px] w-full rounded-[5px] border-2 border-inno-gray bg-white">
              <input
                type="text"
                value={formData.fullName}
                onChange={(e) => handleInputChange("fullName", e.target.value)}
                className="w-full text-inno-gray font-montserrat text-lg sm:text-xl font-normal bg-transparent border-none outline-none placeholder-inno-gray"
                placeholder="Ahmed Baha Eddine Alimi"
              />
            </div>
          </div>

          {/* Email Field */}
          <div className="flex flex-col justify-center items-center gap-[5px] w-full">
            <label className="w-full text-black font-montserrat text-lg sm:text-xl font-normal">
              Email:
            </label>
            <div className="flex h-10 px-6 py-2 items-center gap-[10px] w-full rounded-[5px] border-2 border-inno-gray bg-white">
              <input
                type="email"
                value={formData.email}
                onChange={(e) => handleInputChange("email", e.target.value)}
                className="w-full text-inno-gray font-montserrat text-lg sm:text-xl font-normal bg-transparent border-none outline-none placeholder-inno-gray"
                placeholder="3llimi69@gmail.com"
              />
            </div>
          </div>

          {/* Telegram and Github Row */}
          <div className="flex flex-col sm:flex-row justify-center items-center gap-[15px] w-full">
            {/* Telegram Field */}
            <div className="flex w-full sm:w-[230px] flex-col justify-center items-start gap-[5px]">
              <label className="w-full text-black font-montserrat text-lg sm:text-xl font-normal">
                Telegram:
              </label>
              <div className="flex h-10 px-6 py-2 items-center gap-[10px] w-full rounded-[5px] border-2 border-inno-gray bg-white">
                <input
                  type="text"
                  value={formData.telegram}
                  onChange={(e) =>
                    handleInputChange("telegram", e.target.value)
                  }
                  className="w-full text-inno-gray font-montserrat text-lg sm:text-xl font-normal bg-transparent border-none outline-none placeholder-inno-gray"
                  placeholder="Alias"
                />
              </div>
            </div>

            {/* Github Field */}
            <div className="flex w-full sm:w-[230px] flex-col justify-center items-start gap-[5px]">
              <label className="w-full text-black font-montserrat text-lg sm:text-xl font-normal">
                Github:
              </label>
              <div className="flex h-10 px-6 py-2 items-center gap-[10px] w-full rounded-[5px] border-2 border-inno-gray bg-white">
                <input
                  type="text"
                  value={formData.github}
                  onChange={(e) => handleInputChange("github", e.target.value)}
                  className="w-full text-inno-gray font-montserrat text-lg sm:text-xl font-normal bg-transparent border-none outline-none placeholder-inno-gray"
                  placeholder="Alias"
                />
              </div>
            </div>
          </div>

          {/* Bio Field */}
          <div className="flex flex-col items-start gap-[5px] w-full">
            <label className="w-full text-black font-montserrat text-lg sm:text-xl font-normal">
              Bio:
            </label>
            <div className="flex h-[119px] p-2 items-start gap-[10px] w-full rounded-[5px] border-2 border-inno-green bg-white">
              <textarea
                value={formData.bio}
                onChange={(e) => handleInputChange("bio", e.target.value)}
                className="w-full h-full text-inno-gray font-montserrat text-sm font-normal bg-transparent border-none outline-none resize-none placeholder-inno-gray"
                placeholder="Bio"
              />
            </div>
          </div>
        </div>

        {/* Next Button */}
        <div className="flex w-[100px] h-10 px-[26px] py-2 justify-center items-center gap-[10px] flex-shrink-0 rounded-[10px] bg-inno-green absolute right-4 lg:right-[50px] bottom-4 lg:bottom-[56px] cursor-pointer hover:bg-inno-green/90 transition-colors">
          <span className="text-white font-montserrat text-lg lg:text-xl font-medium">
            Next
          </span>
          <ArrowRight className="w-5 h-5 text-white" />
        </div>
      </div>
    </div>
  );
};

export default ProfileCreationForm;
