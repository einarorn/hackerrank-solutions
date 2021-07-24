#include <iostream>
#include <bits/stdc++.h>

using namespace std;

vector<vector<string>> setupArray();

vector<int> contacts(vector<vector<string>> queries) 
{
    map<string, int> queriesIndex;
    vector<int> results;
    
    for (auto& item : queries)
    {
        if (item[0] == "add")
        {
            for (int i = 1; i <= item[1].length(); i++)
            {
                queriesIndex[item[1].substr(0, i)] += 1;
            }
        } 
        else
        {
            results.push_back(queriesIndex[item[1]]);
        }
    }
    
    return results;
}

int main()
{
    vector<vector<string>> queries = setupArray();
    vector<int> results = contacts(queries);
    
    for (auto& value : results)
    {
        cout<<value;
        cout<<"\n";
    }

    return 0;
}

vector<vector<string>> setupArray() 
{
    vector<vector<string>> queries(4);
    queries[0].resize(2);
    queries[0][0] = "add";
    queries[0][1] = "hack";
    
    queries[1].resize(2);
    queries[1][0] = "add";
    queries[1][1] = "hackerrank";
    
    queries[2].resize(2);
    queries[2][0] = "find";
    queries[2][1] = "hac";
    
    queries[3].resize(2);
    queries[3][0] = "find";
    queries[3][1] = "hak";

    return queries;
}