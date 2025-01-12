#include <vector>
#include <string>
#include <iostream>
#include <locale>
#include <codecvt>

#include "httplib.h"
#include "json.hpp"

using json = nlohmann::json;

extern "C" __declspec(dllexport) std::string SendDetectRequest(const std::string url, const std::vector<std::vector<uint8_t>>& pixelData, const wchar_t* rightKey);

std::string SendDetectRequest(const std::string url, const std::vector<std::vector<uint8_t>>& pixelData, const wchar_t* rightKey) {
    json j;
    json j_pixel = json::array();
    for (const auto& row : pixelData) {
        j_pixel.push_back(row);
    }
    j["pixel"] = j_pixel;
    std::wstring_convert<std::codecvt_utf8<wchar_t>> strCnv;
    j["right_key"] = strCnv.to_bytes(rightKey);

    std::string jsonString = j.dump();

    httplib::Client cli(url);

    httplib::Headers headers = {
        { "ApiKey", "123e4567-e89b-12d3-a456-426614174000" }
    };

    auto res = cli.Post("/", headers, jsonString, "application/json");

    std::string final_result = "";
    if (res && res->status == 200) {
        final_result = res->body;
    }
    else {
        std::cerr << "Request failed: " << (res ? res->body : "No response") << std::endl;
    }
    return final_result;
}